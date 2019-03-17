package main

import (
	"log"
	"my-archive/backend"
	"my-archive/backend/api/handlers"
	appConf "my-archive/backend/config"
	"my-archive/backend/data"
	"my-archive/backend/internal/config"
	am "my-archive/backend/internal/middleware"
	"my-archive/backend/internal/utility"
	casbinadaptor "my-archive/backend/pkg/casbin-adapter"
	"my-archive/backend/usecase"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validator "gopkg.in/go-playground/validator.v8"
)

func main() {
	log.Printf("[INFO] Starting up archive backend service")
	cfg := appConf.SetConfType()
	config.Set(cfg)

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("formida_name", utility.FormidaName)
	}

	repo, err := data.NewPSQL(config.Get("sqlcon-main"))
	if err != nil {
		log.Fatalf("[ERROR] Could not connect to DB: %s", err.Error())
	}

	//e := casbin.NewEnforcer(config.Get("policy-model"), pgAdd.NewMysqlAdapter("", config.Get("mysqlcon")))
	e := casbin.NewEnforcer(config.Get("policy-model"), casbinadaptor.NewPsqlAdapter("", config.Get("sqlcon-policy")))
	e.EnableAutoSave(true)

	s := usecase.NewUseCase(repo, "Europe/Stockholm", e)
	backend.SetUC(s)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsCfg))

	aum, err := am.AuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", aum.LoginHandler)

	//My formida handlers
	adm := r.Group("/admin")
	adm.Use(aum.MiddlewareFunc())
	adm.Use(am.PathAuthorizer(e))
	{
		adm.POST("/create-account", am.ResourceAuthorizer(e, "id", "create-account", "write"), handlers.CreateAccount())
	}

	usr := r.Group("/user")
	usr.Use(aum.MiddlewareFunc())
	usr.Use(am.PathAuthorizer(e))
	{
		usr.POST("/upload", handlers.Upload())
	}
	log.Printf("[INFO] service started at %s", cfg.Get("server"))
	if err := http.ListenAndServe(cfg.Get("server"), r); err != nil {
		log.Fatal(err)
	}
}

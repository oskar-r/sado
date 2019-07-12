package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"my-archive/backend"
	"my-archive/backend/api/handlers"
	"my-archive/backend/api/wshandlers"
	appConf "my-archive/backend/config"
	"my-archive/backend/data"
	"my-archive/backend/internal/config"
	am "my-archive/backend/internal/middleware"
	casbinadaptor "my-archive/backend/pkg/casbin-adapter"
	"my-archive/backend/usecase"
	"net/http"
	"time"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	nats "github.com/nats-io/go-nats"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	log.Printf("[INFO] Starting up archive backend service")
	log.SetFlags(log.Llongfile)
	cfg := appConf.SetConfType()
	config.Set(cfg)

	//Default password is test123
	defaultPwd, err := bcrypt.GenerateFromPassword([]byte("test123"), 12)
	if err != nil {
		log.Fatalf("[ERROR] Could not generate default password %s", err.Error())
	}
	adminPwdPtr := flag.String("admin-password", string(defaultPwd), "Set admin password as a bEncrypted password")
	flag.Parse()
	var adminPwd []byte
	if adminPwdPtr != nil {
		adminPwd = []byte(*(adminPwdPtr))
	} else {
		log.Fatalf("[ERROR] Could parsing password")
	}

	repo, err := data.NewPSQL(config.Get("sqlcon-main"))
	if err != nil {
		log.Fatalf("[ERROR] Could not connect to DB: %s", err.Error())
	}
	err = repo.Ping()
	if err != nil {
		log.Fatalf("[ERROR] Can´t ping DB %s", err.Error())
	}

	repo.ChangeAdminPwd(context.Background(), adminPwd)

	//e := casbin.NewEnforcer(config.Get("policy-model"), pgAdd.NewMysqlAdapter("", config.Get("mysqlcon")))
	e := casbin.NewEnforcer(config.Get("policy-model"), casbinadaptor.NewPsqlAdapter("", config.Get("sqlcon-policy")))
	e.EnableAutoSave(true)

	s := usecase.NewUseCase(repo, "Europe/Stockholm", e)
	backend.SetUC(s)

	httpRoutes, err := setupHTTPServer(e)

	wsRoutes, err := setupWebsockets()

	if err = messageConnection(); err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	}

	if err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	}

	//Start http server
	go func() {
		log.Printf("[INFO] service started at %s", cfg.Get("server"))
		log.Fatal(http.ListenAndServe(cfg.Get("server"), httpRoutes))
	}()

	log.Printf("[INFO] websocket service started at %s", cfg.Get("ws-server"))
	log.Fatal(http.ListenAndServe(cfg.Get("ws-server"), wsRoutes))
}

func setupHTTPServer(e *casbin.Enforcer) (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsCfg))

	aum, err := am.AuthMiddleware()
	if err != nil {
		return nil, errors.New("JWT Error:" + err.Error())
	}

	r.POST("/login", aum.LoginHandler)
	r.GET("/health", handlers.Health())

	//My formida handlers
	adm := r.Group("/admin")
	adm.Use(aum.MiddlewareFunc())
	adm.Use(am.PathAuthorizer(e))
	{
		adm.POST("/create-account", am.ActionAuthorizer(e, "*", "create-account", "write"), handlers.CreateAccount())
	}

	usr := r.Group("/user")
	usr.Use(aum.MiddlewareFunc())
	usr.Use(am.PathAuthorizer(e))
	{
		usr.GET("/config", handlers.Config())
		usr.POST("/changerole", handlers.ChangeRole())
	}

	grp := r.Group("/datasets")
	grp.Use(aum.MiddlewareFunc())
	grp.Use(am.PathAuthorizer(e))
	{
		grp.POST("/upload", handlers.Upload())
		grp.POST("/query", handlers.Query())
		grp.GET("/preview", handlers.Preview())
		grp.GET("/list", handlers.ListMyData())
	}

	return r, nil
}

func setupWebsockets() (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", wshandlers.ServeWS())

	return r, nil
}

func messageConnection() error {
	// Connect Options.
	opts := []nats.Option{nats.Name("Test Subscriber")}
	opts = setupConnOptions(opts)

	opts = append(opts, nats.UserInfo(config.Get("nats-user"), config.Get("nats-pwd")))

	nc, err := nats.Connect(config.Get("nats-server"), opts...)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Connected to nats: %s", config.Get("nats-server"))
	nc.Subscribe("minio-events", func(m *nats.Msg) {
		log.Printf("[DEBUG] %+v", string(m.Data))
		err = backend.ForwardMessage(m.Data)
		if err != nil && err.Error() != "session not found" {
			log.Printf("[ERROR] could not forward message %s", err.Error())
		}
	})
	return nil
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectHandler(func(nc *nats.Conn) {
		log.Printf("Disconnected: will attempt reconnects for %.0fm", totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatal("Exiting, no servers available")
	}))
	return opts
}

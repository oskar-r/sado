package middleware

import (
	"errors"
	"my-archive/backend"
	"my-archive/backend/internal/config"
	jwt "my-archive/backend/internal/gin-jwt"
	"my-archive/backend/internal/utility"
	"my-archive/backend/models" //!!!! Remove dependenciy
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var identityKey = "user"

func AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	if config.Get("rsapriv") == "" {
		return nil, jwt.ErrNoPrivKeyFile
	}
	if config.Get("rsapub") == "" {
		return nil, jwt.ErrNoPubKeyFile
	}

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "archive zone",
		Key:              []byte("thisisnotusedadfarasiknowsinceitisusingRS256asym"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityKey:      identityKey,
		SigningAlgorithm: "RS256",
		PrivKeyFile:      config.Get("rsapriv"),
		PubKeyFile:       config.Get("rsapub"),
		SendCookie:       true,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"username": v.Username,
					"role":     v.CurrentRole,
					"user_id":  v.UserID,
				}
			}
			return jwt.MapClaims{}
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time, data interface{}) {
			var username string
			var role string
			var userID string
			if v, ok := data.(*models.User); ok {
				username = v.Username
				role = v.CurrentRole
				userID = v.UserID
			}

			c.JSON(http.StatusOK, gin.H{
				"code":     http.StatusOK,
				"token":    token,
				"expire":   expire.Format(time.RFC3339),
				"username": username,
				"role":     role,
				"user_id":  userID,
			})
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			if claims["username"] == nil || claims["role"] == nil || claims["user_id"] == nil {
				return &models.User{}
			}
			return &models.User{
				Username:    claims["username"].(string),
				CurrentRole: claims["role"].(string),
				UserID:      claims["user_id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.AuthReqest
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			u, err := backend.ValidateCredentials(loginVals.Username, loginVals.Password)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			if u.UserID == "" {
				return nil, jwt.ErrFailedAuthentication
			}
			if !utility.Contains(loginVals.Role, u.Roles) {
				return nil, errors.New("User does not have the role")
			}
			u.CurrentRole = loginVals.Role
			return u, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

}

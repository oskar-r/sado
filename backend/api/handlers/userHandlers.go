package handlers

import (
	"my-archive/backend"
	"my-archive/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := models.NewUser{}
		err := c.Bind(&u)
		if err != nil {
			resp := map[string]string{"error": err.Error()}
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}
		err = backend.SetUpUserAccount(&u)

		if err != nil {
			resp := map[string]string{"error": err.Error()}
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}
		//map from new user to user

		c.JSON(http.StatusOK, models.User{
			Username: u.Username,
			UserID:   u.UserID,
			MyBucket: u.MyBucket,
		})
		c.Done()
		return
	}
}

func Config() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := extractUserFromClaim(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}
		cfg, err := backend.GetAppConfig(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, cfg)
		c.Done()
		return
	}
}

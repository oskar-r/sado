package handlers

import (
	"encoding/base64"
	"errors"
	"log"
	"my-archive/backend"
	"my-archive/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Query() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := extractUserFromClaim(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}
		q := models.Query{}
		err = c.Bind(&q)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}

		err = backend.QueryBucket(user, &q, c.Request, c.Writer)
		if err != nil {
			log.Printf("[ERROR] %+v", err)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}
		c.Done()
	}
}

func Preview() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := extractUserFromClaim(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}
		file, ok := c.GetQuery("file")
		if !ok {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "base64 encoded filename missing"})
			c.Abort()
			return
		}

		f, err := base64.StdEncoding.DecodeString(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}

		err = backend.PreviewFile(user, f, c.Writer)
		if err != nil {
			log.Printf("[ERROR] %+v", err)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}
		c.Done()
	}
}
func ListMyData() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := extractUserFromClaim(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}

		ds, err := backend.ListMyData(user, c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}

		c.JSON(200, ds)
		c.Done()
		return
	}
}

func extractUserFromClaim(c *gin.Context) (*models.User, error) {
	user := &models.User{}
	if u, ok := c.Get("user"); ok {
		user = u.(*models.User)
	} else {
		return nil, errors.New("no proper claim set")
	}
	if user.Username == "" {
		return nil, errors.New("no username")
	}
	return user, nil
}

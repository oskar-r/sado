package handlers

import (
	"my-archive/backend"
	"my-archive/backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &models.User{}
		if u, ok := c.Get("user"); ok {
			user = u.(*models.User)
		} else {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "no proper claim set"})
			c.Abort()
			return
		}
		if user.Username == "" {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "no username"})
			c.Abort()
			return
		}
		i, err := backend.UploadFile(user, c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}
		type Test struct {
			msg string
		}

		c.JSON(200, &Test{`{"bytes":` + strconv.Itoa(int(i)) + `}`})
		c.Done()
		return
	}
}

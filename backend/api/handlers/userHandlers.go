package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		g, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			resp := map[string]string{"error": err.Error()}
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}
		fmt.Printf("%+v\n", g)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		g, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			resp := map[string]string{"error": err.Error()}
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}
		fmt.Printf("%+v\n", g)
	}
}

package handlers

import (
	"log"
	"my-archive/backend"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := backend.HealthCheck()
		if err != nil {
			log.Printf("[ERROR] Health check: %+v", err)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "Service not healthy"})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})

		c.Done()
		return
	}
}

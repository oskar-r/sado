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

		// key := config.Get("cipher-key")

		/*pwd, err := utility.Encrypt([]byte(key), "testtest")
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			c.Abort()
			return
		}
		log.Printf("[TEST] %s", pwd)
		pwd, err = utility.Decrypt([]byte(key), pwd)
		log.Printf("[TEST] %s", pwd)*/

		// endpoint := "minio.roman.nu"
		//accessKeyID := "oskar"
		//secretAccessKey := "z2yByK2hB1ssIdddJtt3uql@l2gx%W5a"
		// accessKeyID := "test"
		// secretAccessKey := pwd
		// useSSL := true

		// Initialize minio client object.
		// mc, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
		/*if err != nil {
			log.Fatalln(err)
		}

		fs, err := strconv.Atoi(c.GetHeader("Content-length"))
		if err != nil || fs == 0 {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "can't upload empty file"})
			c.Abort()
			return
		}
		if c.Query("name") == "" {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "name query parameter must be set"})
			c.Abort()
			return
		}

		n2, err := mc.PutObject("test2", c.Query("name"), c.Request.Body, int64(fs), minio.PutObjectOptions{
			UserMetadata: map[string]string{
				"user":    user.Username,
				"user-id": user.UserID,
			},
		})
		log.Printf("Size: %+v\n", n2)
		*/

	}
}

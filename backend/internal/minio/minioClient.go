package minio

import (
	"log"
	"my-archive/backend/internal/config"
	"strconv"

	"github.com/minio/minio-go"
)

//mc admin user add roman/my_bucket f5547634-2c08-4780-bda3-e446771c6a2c testtest readwrite
func mainMionio() {
	endpoint := "minio.roman.nu"
	accessKeyID := "oskar"
	secretAccessKey := "z2yByK2hB1ssIdddJtt3uql@l2gx%W5a"
	useSSL := true

	// Initialize minio client object.
	mc, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	policy, err := mc.GetBucketPolicy("test2")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", policy) // minioClient is now setup

}

func ConnectToMinio(userID, creds string) (*minio.Client, error) {
	https, _ := strconv.ParseBool(config.Get("minio-https"))
	mc, err := minio.New(config.Get("minio-server"), userID, creds, https)
	if err != nil {
		log.Printf("[ERROR] Creating client: %+v", err)
		return nil, err
	}
	return mc, err
}

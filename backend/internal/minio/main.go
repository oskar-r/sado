package minio

import (
	"log"

	"github.com/minio/minio-go"
)

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

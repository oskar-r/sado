package minio

import (
	"errors"
	"io"
	"log"

	minio "github.com/minio/minio-go"
)

type Configuration struct {
	AccessKeyID string
	SecretKey   string
	Filename    string
	BucketName  string
	Filesize    int64

	File     io.Reader
	Metadata map[string]string
	Endpoint string
	UseSSL   bool
}

func Upload(conf *Configuration) (int64, error) {

	// Initialize minio client object.

	mc, err := minio.New(conf.Endpoint, conf.AccessKeyID, conf.SecretKey, conf.UseSSL)
	if err != nil {
		log.Printf("[ERROR] Creating client: %+v", err)
		return 0, err
	}

	lb, err := mc.BucketExists(conf.BucketName)
	if err != nil {
		log.Printf("[ERROR] Could not check for bucket existance: %+v", err)
		return 0, err
	}
	if !lb {
		log.Printf("[ERROR] Bucket %s does not exist", conf.BucketName)
		return 0, errors.New("bucket does not exits")
	}

	fs, err := mc.PutObject(conf.BucketName, conf.Filename, conf.File, conf.Filesize, minio.PutObjectOptions{
		UserMetadata: conf.Metadata,
		ContentType:  conf.Metadata["content_type"],
	})
	if err != nil {
		log.Printf("[ERROR] Upload object: %+v", err)
	}

	return fs, err
}

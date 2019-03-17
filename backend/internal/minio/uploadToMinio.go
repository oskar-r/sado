package minio

import (
	"io"

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
		return 0, err
	}

	fs, err := mc.PutObject(conf.BucketName, conf.Filename, conf.File, conf.Filesize, minio.PutObjectOptions{
		UserMetadata: conf.Metadata,
	})
	return fs, nil
}

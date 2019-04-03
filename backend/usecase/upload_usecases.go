package usecase

import (
	"encoding/base64"
	"errors"
	"my-archive/backend/internal/config"
	"my-archive/backend/internal/minio"
	"my-archive/backend/models"
	"net/http"
	"strconv"
)

func (uc *usecase) UploadFile(user *models.User, request *http.Request) (int64, error) {
	creds, err := getCredentials(uc, user)

	if err != nil {
		return 0, err
	}

	if request.ContentLength == 0 {
		return 0, errors.New("trying to upload empty filet")
	}
	if request.URL.Query().Get("name") == "" {
		return 0, errors.New("no file name provided")
	}
	filename, err := base64.StdEncoding.DecodeString(request.URL.Query().Get("name"))
	if err != nil {
		return 0, errors.New("filename not properly encoded")
	}
	contentType, err := base64.StdEncoding.DecodeString(request.URL.Query().Get("content-type"))
	if err != nil {
		return 0, errors.New("content type not set")
	}
	https, err := strconv.ParseBool(config.Get("minio-https"))
	if err != nil {
		https = true
	}

	fs, err := minio.Upload(&minio.Configuration{
		AccessKeyID: user.UserID,
		SecretKey:   creds.password,
		BucketName:  creds.bucket,
		Filename:    string(filename),
		Filesize:    request.ContentLength,
		File:        request.Body,
		Metadata: map[string]string{
			"username":     user.Username,
			"user_id":      user.UserID,
			"content_type": string(contentType),
		},
		Endpoint: config.Get("minio-server"),
		UseSSL:   https,
	})
	request.Body.Close()

	return fs, err
}

package usecase

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"my-archive/backend/internal/config"
	"my-archive/backend/internal/minio"
	"my-archive/backend/internal/utility"
	"my-archive/backend/models"
	"net/http"
)

func (uc *usecase) UploadFile(user *models.User, request *http.Request) (int64, error) {
	if user.UserID == "" {
		return 0, errors.New("bad request")
	}

	sc, err := uc.repo.GetStorageCredentials(context.Background(), user.UserID)
	if err != nil {
		return 0, err
	}

	var creds string
	if val, ok := sc["secret"]; ok {
		creds, err = utility.Decrypt([]byte(config.Get("cipher-key")), val)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("User lack bucket credentials")
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

	log.Printf("[TEST] %s", creds)

	fs, err := minio.Upload(&minio.Configuration{
		AccessKeyID: user.UserID,
		SecretKey:   creds,
		BucketName:  sc["bucket"],
		Filename:    string(filename),
		Filesize:    request.ContentLength,
		File:        request.Body,
		Metadata: map[string]string{
			"username":     user.Username,
			"user_id":      user.UserID,
			"content_type": string(contentType),
		},
		Endpoint: "minio.roman.nu",
		UseSSL:   true,
	})
	request.Body.Close()

	return fs, err
}

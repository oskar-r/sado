package usecase

import (
	"context"
	"errors"
	"my-archive/backend/internal/config"
	"my-archive/backend/internal/minio"
	"my-archive/backend/internal/utility"
	"my-archive/backend/models"
	"net/http"
)

func (uc *usecase) UploadFile(user *models.User, request *http.Request) (int, error) {
	sc, err := uc.repo.GetStorageCredentials(context.Background(), user.Username)
	if err != nil {
		return 0, err
	}
	var creds string
	if val, ok := sc["credentials"]; ok {
		creds, err = utility.Decrypt([]byte(config.Get("cipher-key")), val)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("bucket credentials not set")
	}

	if request.ContentLength == 0 {
		return 0, errors.New("trying to upload empty filet")
	}
	if request.URL.Query().Get("name") == "" {
		return 0, errors.New("no file name provided")
	}

	minio.Upload(&minio.Configuration{
		AccessKeyID: user.UserID,
		SecretKey:   creds,
		BucketName:  sc["bucket"],
		Filename:    request.URL.Query().Get("name"),
		Filesize:    request.ContentLength,
		File:        request.Body,
		Metadata: map[string]string{
			"username": user.Username,
			"user_id":  user.UserID,
		},
		Endpoint: "storage.roman.nu",
		UseSSL:   true,
	})
	request.Body.Close()
	
	return 0, nil
}

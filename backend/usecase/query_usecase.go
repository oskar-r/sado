package usecase

import (
	"context"
	"errors"
	"io"
	"my-archive/backend/internal/config"
	"my-archive/backend/internal/minio"
	"my-archive/backend/internal/utility"
	"my-archive/backend/models"
	"net/http"
	"strings"
)

func (uc *usecase) PreviewFile(user *models.User, filename []byte, pipe io.Writer) error {
	sc, err := getCredentials(uc, user)
	if err != nil {
		return err
	}

	err = minio.Preview(pipe, user.UserID, sc.password, sc.bucket, string(filename))
	if err != nil {
		return err
	}
	return nil
}

func (uc *usecase) QueryBucket(user *models.User, query *models.Query, request *http.Request, pipe io.Writer) error {
	sc, err := getCredentials(uc, user)
	if err != nil {
		return err
	}

	query.Query = strings.Replace(query.Query, query.Dataset, "s3Object", 1)

	err = minio.Query(pipe, user.UserID, sc.password, sc.bucket, query)
	if err != nil {
		return err
	}
	return nil
}

func (uc *usecase) ListMyData(user *models.User, request *http.Request) ([]models.DataSet, error) {
	sc, err := getCredentials(uc, user)
	if err != nil {
		return nil, err
	}
	ds, err := minio.ListDatasetsInBucket(user.UserID, sc.password, sc.bucket)
	if err != nil {
		return nil, err
	}
	return ds, nil
}

type credentials struct {
	password string
	bucket   string
}

func getCredentials(uc *usecase, user *models.User) (*credentials, error) {

	if user.UserID == "" {
		return nil, errors.New("bad request")
	}
	sc, err := uc.repo.GetStorageCredentials(context.Background(), user.UserID)
	if err != nil {
		return nil, err
	}

	var creds string
	if val, ok := sc["secret"]; ok {
		creds, err = utility.Decrypt([]byte(config.Get("cipher-key")), val)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("User lack bucket credentials")
	}
	if creds == "" || sc["bucket"] == "" {
		return nil, errors.New("No bucket or storage password")
	}
	return &credentials{
		password: creds,
		bucket:   sc["bucket"],
	}, nil
}

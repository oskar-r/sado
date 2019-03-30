package backend

import (
	"io"
	"my-archive/backend/models"
	"net/http"
)

//UseCase use case interface
type UploadUseCase interface {
	UploadFile(user *models.User, req *http.Request) (int64, error)
	QueryBucket(user *models.User, query *models.Query, request *http.Request, pipe io.Writer) error
	ListMyData(user *models.User, request *http.Request) ([]models.DataSet, error)
	PreviewFile(user *models.User, filename []byte, pipe io.Writer) error
}

func UploadFile(user *models.User, req *http.Request) (int64, error) {
	return impl.UploadFile(user, req)
}
func QueryBucket(user *models.User, query *models.Query, request *http.Request, pipe io.Writer) error {
	return impl.QueryBucket(user, query, request, pipe)
}

func ListMyData(user *models.User, request *http.Request) ([]models.DataSet, error) {
	return impl.ListMyData(user, request)
}

func PreviewFile(user *models.User, filename []byte, pipe io.Writer) error {
	return impl.PreviewFile(user, filename, pipe)
}

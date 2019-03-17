package backend

import (
	"my-archive/backend/models"
	"net/http"
)

//UseCase use case interface
type UploadUseCase interface {
	UploadFile(user *models.User, req *http.Request) (int, error)
}

func UploadFile(user *models.User, req *http.Request) (int, error) {
	return impl.UploadFile(user, req)
}

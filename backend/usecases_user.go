package backend

import (
	"my-archive/backend/models"
	"time"
)

//UseCase use case interface
type UserUseCase interface {
	ValidateCredentials(username string, password string) (*models.User, error)
	CreateUserSession(userID int) (string, error)
	DeleteUserSession(userID int) (string, error)
	LoginAtempt(userID int, ip string, loginAt *time.Time, success bool) error
	TokenRefresh(userID int, role string) error
	SetUpUserAccount(user *models.NewUser) error
	GetAppConfig(user *models.User) (*models.AppConfig, error)
}

func ValidateCredentials(username string, password string) (*models.User, error) {
	return impl.ValidateCredentials(username, password)
}

func SetUpUserAccount(user *models.NewUser) error {
	return impl.SetUpUserAccount(user)
}

func GetAppConfig(user *models.User) (*models.AppConfig, error) {
	return impl.GetAppConfig(user)
}
func CreateUserSession(userID int) (string, error) {
	return impl.CreateUserSession(userID)
}

func LoginAtempt(userID int, ip string, loginAt *time.Time, success bool) error {
	return impl.LoginAtempt(userID, ip, loginAt, success)
}

func TokenRefresh(userID int, role string) error {
	return impl.TokenRefresh(userID, role)
}

package usecase

import (
	"context"
	"errors"
	"log"
	"my-archive/backend/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//ValidateCredentials validate a user based on username and pipeline. Returns a pointer to the user object or an error
func (uc *usecase) ValidateCredentials(username string, password string) (*models.User, error) {
	u, err := uc.repo.GetUser(context.Background(), username, "")
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	if err != nil {
		return nil, errors.New("Bad password/username combo")
	}
	u.Roles = uc.enforcer.GetRolesForUser(u.UserID)
	log.Printf("[TEST] %+v", u.Roles)

	return &u, nil
}

func (uc *usecase) CreateUserSession(userID int) (string, error) {
	panic("not implemented")
}
func (uc *usecase) DeleteUserSession(userID int) (string, error) {
	panic("not implemented")
}
func (uc *usecase) LoginAtempt(userID int, ip string, loginAt *time.Time, success bool) error {
	panic("not implemented")
}
func (uc *usecase) TokenRefresh(userID int, role string) error {
	panic("not implemented")
}

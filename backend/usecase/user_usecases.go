package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"my-archive/backend/internal/config"
	"my-archive/backend/internal/minio"
	"my-archive/backend/internal/utility"
	"my-archive/backend/models"
	"time"

	uuid "github.com/satori/go.uuid"
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
	if len(u.Roles) == 0 { //Default role is user
		u.Roles = append(u.Roles, "user")
	}

	log.Printf("[TEST] %+v", u.Roles)

	return &u, nil
}

func (uc *usecase) SetUpUserAccount(user *models.NewUser) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {

	}
	user.BCryptPassword = string(pwd)

	str, err := utility.Encrypt([]byte(config.Get("cipher-key")), user.Password)
	if err != nil {
		log.Printf("[ERROR] %+v", err)
		return err
	}
	user.AESPassword = str

	uuid := uuid.NewV4()

	user.UserID = uuid.String()

	err = minio.CreateDefaultBucket(user.MyBucket)
	if err != nil {
		log.Printf("[ERROR] %+v", err)
		return err
	}

	err = minio.CreateUser(user.UserID, user.Password, user.MyBucket)
	if err != nil {
		log.Printf("[ERROR] %+v", err)
		return err
	}

	err = uc.repo.CreateUserAccount(context.Background(), user)
	if err != nil {
		err2 := minio.DeleteBucket(user.MyBucket) //Remove bucket of account creation fails
		if err2 != nil {
			log.Printf("[ERROR] %+v", err2)
		}
		log.Printf("[ERROR] %+v", err)
		return err
	}

	return nil
}

func (uc *usecase) GetAppConfig(user *models.User) (*models.AppConfig, error) {
	if user.CurrentRole == "" {
		return nil, errors.New("No active role set")
	}
	config, err := uc.repo.GetAppConfig(context.Background(), user.CurrentRole)
	if err != nil {
		return nil, err
	}

	appConf := &models.AppConfig{}

	//Make sh
	err = json.Unmarshal([]byte(config), &appConf)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}
	appConf.ActiveRole = user.CurrentRole
	appConf.Roles = uc.enforcer.GetRolesForUser(user.UserID)
	if len(appConf.Roles) == 0 {
		appConf.Roles = append(appConf.Roles, user.CurrentRole)
	}
	return appConf, err
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

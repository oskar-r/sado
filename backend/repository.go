package backend

import (
	"context"
	"my-archive/backend/models"

	"github.com/casbin/casbin/model"
	"github.com/gorilla/websocket"
)

type Utility interface {
	Close()
	Ping() error
}

type AuthzRepository interface {
	LoadPolicies(namespace string) ([]string, error)
	CreatePolicyDB() error
	SaveAllPolicy(namespace string, model model.Model) error
	AddPolicy(i []interface{}) (int64, error)
	DeletePolicy(i []interface{}) (int64, error)
}
type UserRepository interface {
	GetUser(ctx context.Context, username string, userID string) (models.User, error)
	GetStorageCredentials(ctx context.Context, userID string) (map[string]string, error)
	CreateUserAccount(ctx context.Context, user *models.NewUser) error
	GetAppConfig(ctx context.Context, role string) (string, error)
	ChangeAdminPwd(ctx context.Context, bCryptPwd []byte) error
}

type SessionStore interface {
	SaveSession(id string, conn *websocket.Conn) error
	DeleteSession(id string) error
	GetSession(id string) (*websocket.Conn, error)
}

type Repository interface {
	AuthzRepository
	UserRepository
	Utility
	SessionStore
}

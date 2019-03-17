package data

import (
	"context"
	"fmt"
	"log"
	"my-archive/backend/models"
)

const userSchema = "public.users"

const userAndRolesQ = `SELECT ` + userSchema + `.username, ` + userSchema + `.user_id, ` + userSchema + `.user_pass FROM ` + userSchema + ` WHERE %s = $1 `

func (r *PSQLRepo) GetUser(ctx context.Context, username string, userID string) (models.User, error) {
	u := []models.User{}
	um := models.User{}

	var err error
	var q string
	if userID != "" {
		q = fmt.Sprintf(userAndRolesQ, userSchema+".user_id")
		err = r.db.Select(&u, q, userID)
	} else {
		q = fmt.Sprintf(userAndRolesQ, userSchema+".username")
		err = r.db.Select(&u, q, username)
	}

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return um, err
	}
	var roles []string

	for _, v := range u {
		roles = append(roles, v.CurrentRole)
	}
	um = u[0]
	um.Roles = roles
	um.Username = username

	return um, err
}

const storageSchema = "public.minio"
const storageCredentialsQ = `SELECT ` + storageSchema + `.bucket_name, ` + storageSchema + `.bucket_credentials FROM ` + storageSchema + ` WHERE ` + storageSchema + `.user_id = $1 `

func (r *PSQLRepo) GetStorageCredentials(ctx context.Context, userID string) (map[string]string, error) {

	type TM struct {
		Bucket      string
		Credentials string
	}
	tm := TM{}
	err := r.db.Select(&tm, storageCredentialsQ, userID)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	model := map[string]string{
		"bucket":      tm.Bucket,
		"credentials": tm.Credentials,
	}

	return model, err
}

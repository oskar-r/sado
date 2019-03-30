package data

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"my-archive/backend/internal/utility"
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
const storageCredentialsQ = `SELECT ` + storageSchema + `.bucket_name, 
	` + storageSchema + `.bucket_access_key, 
	` + storageSchema + `.bucket_secret 
	FROM ` + storageSchema +
	` WHERE ` + storageSchema + `.user_id = $1 `

func (r *PSQLRepo) GetStorageCredentials(ctx context.Context, userID string) (map[string]string, error) {

	type TM struct {
		BucketName      string `db:"bucket_name"`
		BucketAccessKey string `db:"bucket_access_key"`
		BucketSecret    string `db:"bucket_secret"`
	}
	tm := TM{}
	log.Printf("[TEST] %s", userID)
	err := r.db.Get(&tm, storageCredentialsQ, userID)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	model := map[string]string{
		"bucket":     tm.BucketName,
		"access-key": tm.BucketAccessKey,
		"secret":     tm.BucketSecret,
	}

	return model, err
}

func (r *PSQLRepo) CreateUserAccount(ctx context.Context, user *models.NewUser) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}
	err = createAccount(tx, user)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		err2 := tx.Rollback()
		if err2 != nil {
			log.Printf("[ERROR] %s", err2.Error())
		}
		return err
	}
	err = createBucket(tx, user)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			log.Printf("[ERROR] %s", err2.Error())
		}
		return err
	}
	return tx.Commit()
}

const createUserAccountQ = `INSERT INTO users(
	username, 
	user_id,
	user_pass) VALUES($1,$2,$3)`

func createAccount(tx *sql.Tx, user *models.NewUser) error {
	res, err := tx.Exec(createUserAccountQ, user.Username, user.UserID, user.BCryptPassword)
	return utility.ExecReturn(res, err)
}

const createUserBucketQ = `INSERT INTO minio(
	user_id, 
	bucket_name,
	bucket_secret, 
	bucket_access_key) VALUES($1,$2,$3,$4)`

func createBucket(tx *sql.Tx, user *models.NewUser) error {
	res, err := tx.Exec(createUserBucketQ, user.UserID, user.MyBucket, user.AESPassword, user.UserID)
	return utility.ExecReturn(res, err)
}

const getAppConfig = `SELECT configuration FROM app_config WHERE role = $1`

func (r *PSQLRepo) GetAppConfig(ctx context.Context, role string) (string, error) {
	var cfg string
	err := r.db.Get(&cfg, getAppConfig, role)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}
	return cfg, err
}

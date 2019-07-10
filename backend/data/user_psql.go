package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"my-archive/backend/internal/utility"
	"my-archive/backend/models"
)

const userSchema = "public.users"

const userAndRolesQ = `SELECT ` + userSchema + `.username, ` + userSchema + `.user_id, ` + userSchema + `.user_pass FROM ` + userSchema + ` WHERE %s = $1 `

func (r *Repos) GetUser(ctx context.Context, username string, userID string) (models.User, error) {
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

	if len(u) == 0 {
		log.Printf("[ERROR] User %s does not exits", username)
		return um, errors.New("User does not exist")
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
	` + storageSchema + `.bucket_secret,
	` + storageSchema + `.user_bucket_name
	FROM ` + storageSchema +
	` WHERE ` + storageSchema + `.user_id = $1 `

func (r *Repos) GetStorageCredentials(ctx context.Context, userID string) (map[string]string, error) {

	type TM struct {
		UserBucketName  string `db:"user_bucket_name"`
		BucketName      string `db:"bucket_name"`
		BucketAccessKey string `db:"bucket_access_key"`
		BucketSecret    string `db:"bucket_secret"`
	}
	tm := TM{}

	err := r.db.Get(&tm, storageCredentialsQ, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("No bucket created for user")
		}
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}

	model := map[string]string{
		"bucket":           tm.BucketName,
		"access-key":       tm.BucketAccessKey,
		"secret":           tm.BucketSecret,
		"user-bucket-name": tm.UserBucketName,
	}

	return model, err
}

func (r *Repos) CreateUserAccount(ctx context.Context, user *models.NewUser) error {
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
	user_bucket_name,
	bucket_name,
	bucket_secret, 
	bucket_access_key) VALUES($1,$2,$3,$4,$5)`

func createBucket(tx *sql.Tx, user *models.NewUser) error {
	res, err := tx.Exec(createUserBucketQ, user.UserID, user.MyBucket, user.MinioBucketName, user.AESPassword, user.UserID)
	return utility.ExecReturn(res, err)
}

const getAppConfig = `SELECT configuration FROM app_config WHERE role = $1`

func (r *Repos) GetAppConfig(ctx context.Context, role string) (string, error) {
	var cfg string
	log.Printf("[TEST]: %s", role)

	err := r.db.Get(&cfg, getAppConfig, role)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}
	return cfg, err
}

const changeAdminPwd = `UPDATE users SET user_pass=$1 WHERE username='admin'`

func (r *Repos) ChangeAdminPwd(ctx context.Context, bCryptPwd []byte) error {
	_, err := r.db.Exec(changeAdminPwd, bCryptPwd)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}
	return err
}

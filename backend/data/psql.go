package data

import (
	"fmt"
	"runtime"

	"github.com/jmoiron/sqlx"
)

type PSQLRepo struct {
	db *sqlx.DB
}

func NewPSQL(connectionString string) (*PSQLRepo, error) {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &PSQLRepo{
		db,
	}, nil
}

func (r *PSQLRepo) Close() {
	r.db.Close()
}

func (r *PSQLRepo) Ping() error {
	return r.db.Ping()
}
func Debug() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s %d", file, line)
}

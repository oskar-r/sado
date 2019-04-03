package data

import (
	"fmt"
	"runtime"

	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
)

type Repos struct {
	db           *sqlx.DB
	sessionStore map[string]*websocket.Conn
}

func NewPSQL(connectionString string) (*Repos, error) {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	ss := make(map[string]*websocket.Conn, 0)

	return &Repos{
		db,
		ss,
	}, nil
}

func (r *Repos) Close() {
	r.db.Close()
}

func (r *Repos) Ping() error {
	return r.db.Ping()
}
func Debug() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s %d", file, line)
}

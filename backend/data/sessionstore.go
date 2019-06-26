package data

import (
	"errors"

	"github.com/gorilla/websocket"
)

func (r *Repos) SaveSession(id string, conn *websocket.Conn) error {
	r.sessionStore[id] = conn
	return nil
}

func (r *Repos) DeleteSession(id string) error {
	if _, ok := r.sessionStore[id]; ok {
		delete(r.sessionStore, id)
	}
	return nil
}

func (r *Repos) GetSession(id string) (*websocket.Conn, error) {
	c, ok := r.sessionStore[id]
	if !ok {
		return nil, errors.New("session  " + id + " not found")
	}
	return c, nil
}

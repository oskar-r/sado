package backend

import "github.com/gorilla/websocket"

type SessionManagement interface {
	SaveSession(id string, conn *websocket.Conn) error
	DeleteSession(id string) error
	ForwardMessage(b []byte) error
}

func SaveSession(id string, conn *websocket.Conn) error {
	return impl.SaveSession(id, conn)
}

func DeleteSession(id string) error {
	return impl.DeleteSession(id)
}

func ForwardMessage(b []byte) error {
	return impl.ForwardMessage(b)
}

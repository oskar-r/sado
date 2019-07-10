package wshandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"my-archive/backend"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
)

type socketContext struct {
	store  *sessions.CookieStore
	socket *websocket.Conn
	// ... and the rest of our globals.
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type pushMsg struct {
	User    string `json:"user,omitempty"`
	Message string `json:"message,omitempty"`
}

func supervisor(cm map[string]*websocket.Conn, w http.ResponseWriter, r *http.Request) {
	type supResp struct {
		ActiveUsers []string `json:"active_users"`
	}
	var sr supResp

	for i := range cm {
		sr.ActiveUsers = append(sr.ActiveUsers, i)
	}

	rsp, err := json.Marshal(sr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	fmt.Fprintf(w, string(rsp))
}

//ServeWS creates a Websocket server and listens to incomming requests
func ServeWS() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		id, err := c.Cookie("id")
		if err != nil {
			log.Printf("Could not read cookie: %s\n", err.Error())
			return
		}
		if id != "" {
			if conn == nil {
				log.Printf("nil connection for id:%s\n", id)
				return
			}
			err = backend.SaveSession(id, conn)
			if conn == nil {
				log.Printf("[ERROR] Could not save websocket session")
				return
			}

			if err != nil {
				log.Printf("upgrade:%s\n", err.Error())
				return
			}
			log.Printf("[INFO] Websocket connection made by id: %s", id)
			for {
				mt, message, err := conn.ReadMessage()
				if err != nil {
					log.Println("[ERROR] Read:", err)
					conn.Close()
					backend.DeleteSession(id)
					break
				}
				err = conn.WriteMessage(mt, message)
				if err != nil {
					log.Println("[ERROR] Write:", err)
					conn.Close()
					backend.DeleteSession(id)
					break
				}
			}
			conn.Close()
			backend.DeleteSession(id)
		}
	}
}

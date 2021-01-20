package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const (
	READBUFFERSIZE = 1024
	WRITEBUFFERSIZE = 1024

)

var upgrader = websocket.Upgrader{
	ReadBufferSize: READBUFFERSIZE,
	WriteBufferSize: WRITEBUFFERSIZE,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}



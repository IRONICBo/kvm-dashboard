package ws

import (
	"net/http"

	"github.com/olahol/melody"
)

type WSServer struct {
	server  *melody.Melody
	sessMap map[string]*melody.Session
}

func NewWSServer() *WSServer {
	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return &WSServer{
		server:  m,
		sessMap: make(map[string]*melody.Session),
	}
}

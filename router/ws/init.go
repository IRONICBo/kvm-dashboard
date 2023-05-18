package ws

import (
	"errors"
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

func (ws *WSServer) WriteData(uuid string, data string) error {
	sess, ok := ws.sessMap[uuid]
	if !ok {
		return errors.New("can not find session")
	}

	sess.Write([]byte(data))

	return nil
}

func (ws *WSServer) CloseSession(uuid string) error {
	sess, ok := ws.sessMap[uuid]
	if !ok {
		return errors.New("can not find session")
	}

	sess.Close()

	return nil
}

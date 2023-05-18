package ws

import (
	"errors"
	"kvm-dashboard/utils"
	"net/http"

	"github.com/olahol/melody"
)

var ControlWSServer *WSServer

func NewControlWSServer() {
	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ControlWSServer = &WSServer{
		server:  m,
		sessMap: make(map[string]*melody.Session),
	}

	// hook
	ControlWSServer.server.HandleConnect(ControlWSServer.HandleControlConnect)
	ControlWSServer.server.HandleClose(ControlWSServer.HandleControlClose)
}

func (ws *WSServer) HandleControlRequest(w http.ResponseWriter, r *http.Request) error {
	err := ws.server.HandleRequest(w, r)
	if err != nil {
		utils.Log.Error("Can not handle request", err)
		return err
	}

	return nil
}

func (ws *WSServer) HandleControlConnect(s *melody.Session) {
	uuid := s.Request.URL.Query().Get("UUID")

	ws.sessMap[uuid] = s
}

func (ws *WSServer) HandleControlClose(s *melody.Session, msg int, reason string) error {
	uuid := s.Request.URL.Query().Get("UUID")

	delete(ws.sessMap, uuid)
	return nil
}

func (ws *WSServer) WriteData(uuid string, data string) error {
	sess, ok := ws.sessMap[uuid]
	if !ok {
		return errors.New("can not find session")
	}

	sess.Write([]byte(data))
	return nil
}

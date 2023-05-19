package ws

import (
	"kvm-dashboard/utils"
	"net/http"

	"github.com/olahol/melody"
)

var SimpleWSServer *WSServer

func NewSimpleWSServer() {
	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	SimpleWSServer = &WSServer{
		server:  m,
		sessMap: make(map[string]*melody.Session),
	}

	// hook
	SimpleWSServer.server.HandleConnect(SimpleWSServer.HandleSimpleConnect)
	SimpleWSServer.server.HandleClose(SimpleWSServer.HandleSimpleClose)
}

func (ws *WSServer) HandleSimpleRequest(w http.ResponseWriter, r *http.Request) error {
	err := ws.server.HandleRequest(w, r)
	if err != nil {
		utils.Log.Error("Can not handle request", err)
		return err
	}

	return nil
}

func (ws *WSServer) HandleSimpleConnect(s *melody.Session) {
	uuid := s.Request.URL.Query().Get("UUID")

	ws.sessMap[uuid] = s
}

func (ws *WSServer) HandleSimpleClose(s *melody.Session, msg int, reason string) error {
	uuid := s.Request.URL.Query().Get("UUID")

	delete(ws.sessMap, uuid)
	return nil
}

package ws

import (
	"errors"
	"kvm-dashboard/consts"
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
		sessMap: make(map[string]map[string]*melody.Session),
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
	// todo: replace with user id
	token := s.Request.URL.Query().Get("token")
	if token == "" {
		token = consts.DEFAULT_TOKEN
	}

	// if not exist, create
	if _, ok := ws.sessMap[uuid]; !ok {
		ws.sessMap[uuid] = make(map[string]*melody.Session)
	}
	ws.sessMap[uuid][token] = s
}

func (ws *WSServer) HandleSimpleClose(s *melody.Session, msg int, reason string) error {
	uuid := s.Request.URL.Query().Get("UUID")
	// todo: replace with user id
	token := s.Request.URL.Query().Get("token")
	if token == "" {
		token = consts.DEFAULT_TOKEN
	}

	sessDevice, ok := ws.sessMap[uuid]
	if !ok {
		return errors.New("can not find session")
	}

	delete(sessDevice, token)

	return nil
}

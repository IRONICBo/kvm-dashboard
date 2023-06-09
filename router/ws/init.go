package ws

import (
	"errors"
	"net/http"

	"github.com/olahol/melody"
)

type WSServer struct {
	server  *melody.Melody
	sessMap map[string]map[string]*melody.Session
}

func NewWSServer() *WSServer {
	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return &WSServer{
		server:  m,
		sessMap: make(map[string]map[string]*melody.Session),
	}
}

// broadcast
func (ws *WSServer) WriteData(uuid string, data string) error {
	sessDevice, ok := ws.sessMap[uuid]
	if !ok {
		return errors.New("can not find session")
	}

	// write to all devices
	for _, sess := range sessDevice {
		sess.Write([]byte(data))
	}

	return nil
}

func (ws *WSServer) WriteDataToDevice(uuid string, device string, data string) error {
	sessDevice, ok := ws.sessMap[uuid]
	if !ok {
		return errors.New("can not find session")
	}

	sess, ok := sessDevice[device]
	if !ok {
		return errors.New("can not find session")
	}

	sess.Write([]byte(data))

	return nil
}

func (ws *WSServer) CloseSessionWithDevice(uuid string, device string) error {
	sessDevice, ok := ws.sessMap[uuid]
	if !ok {
		return errors.New("can not find session")
	}

	sess, ok := sessDevice[device]
	if !ok {
		return errors.New("can not find session")
	}

	sess.Close()

	return nil
}

// boardcast
func (ws *WSServer) CloseSession(uuid string) error {
	sessDevice, ok := ws.sessMap[uuid]
	if !ok {
		return errors.New("can not find session")
	}

	for _, sess := range sessDevice {
		sess.Close()
	}

	return nil
}

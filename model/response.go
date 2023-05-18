package model

import "encoding/json"

type Progress struct {
	UUID    string `json:"uuid"`
	Percent int    `json:"percent"`
	Msg     string `json:"msg"`
	State   bool   `json:"state"` // true: success, false: failed
}

func NewProgress(uuid string, percent int, msg string, state bool) *Progress {
	return &Progress{
		UUID:    uuid,
		Percent: percent,
		Msg:     msg,
		State:   state,
	}
}

func NewProgressJson(uuid string, percent int, msg string, state bool) string {
	progress := NewProgress(uuid, percent, msg, state)
	jsonByte, _ := json.Marshal(progress)
	return string(jsonByte)
}

package model

import (
	"encoding/json"
	"net/http"
)

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

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Success() *Response {
	return NewResponse(http.StatusOK, "success", nil)
}

func SuccessWithMsg(msg string) *Response {
	return NewResponse(http.StatusOK, msg, nil)
}

func SuccessWithData(data interface{}) *Response {
	return NewResponse(http.StatusOK, "success", data)
}

func SuccessWithDataAndMsg(data interface{}, msg string) *Response {
	return NewResponse(http.StatusOK, msg, data)
}

func Failed() *Response {
	return NewResponse(http.StatusInternalServerError, "failed", nil)
}

func FailedWithMsg(msg string) *Response {
	return NewResponse(http.StatusInternalServerError, msg, nil)
}

func FailedWithData(data interface{}) *Response {
	return NewResponse(http.StatusInternalServerError, "failed", data)
}

func FailedWithDataAndMsg(data interface{}, msg string) *Response {
	return NewResponse(http.StatusInternalServerError, msg, data)
}

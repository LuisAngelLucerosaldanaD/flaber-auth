package models

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error bool `json:"error"`
	Data interface{} `json:"data"`
	Code int `json:"code"`
	Type string `json:"type"`
	Msg string `json:"msg"`
}

func (rm *Response) set(err bool, data interface{}, code int)  {
	rm.Error = err
	rm.Data = data
	rm.Code = code
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{})  {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
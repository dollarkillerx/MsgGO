package defs

import "net/http"

type success struct {
	Data string `json:"data"`
	Code string `json:"code"`
}

type SuccessResponse struct {
	HttpSC int
	Data success
}

var (
	SuccessRegisterOK = SuccessResponse{http.StatusCreated,success{"Register success","200"}}
)
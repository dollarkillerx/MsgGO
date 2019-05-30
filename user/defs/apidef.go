package defs

import "net/http"

type Success struct {
	Data string `json:"data"`
	Code string `json:"code"`
}

type SuccessResponse struct {
	HttpSC int
	Data Success
}

var (
	SuccessRegisterOK = SuccessResponse{http.StatusCreated,Success{"Register success","200"}}
)
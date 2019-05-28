package defs

import "net/http"

type err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{http.StatusBadRequest,err{"Request body is not correct","001"}}
	ErrorNotAuthUser = ErrorResponse{http.StatusUnauthorized,err{"User authentication failed","002"}}
	ErrorDBError = ErrorResponse{http.StatusInternalServerError,err{"DB ops failed","003"}}
	ErrorInternalFaults = ErrorResponse{http.StatusInternalServerError,err{"Internal service error","004"}}
)
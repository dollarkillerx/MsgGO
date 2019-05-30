package defs

import "net/http"

type err struct {
	Error string `json:"data"`
	ErrorCode string `json:"code"`
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
	ErrorRegister = ErrorResponse{http.StatusBadRequest,err{"Users have registered !!!","005"}}
	ErrorGenerateToken = ErrorResponse{http.StatusInternalServerError,err{"Generate Token error !!!","006"}}
)
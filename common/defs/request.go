/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-25
* Time: 上午11:42
* */
package defs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestDate struct {
	Code int
	Data interface{}
}

var (
	ErrorBadRequest = &RequestDate{Code: http.StatusBadRequest, Data: &gin.H{"Code": "400", "Message": "Bad Request"}}
	ErrorBanAuth    = &RequestDate{Code: http.StatusUnauthorized, Data: &gin.H{"Code": "401", "Message": "User Unauthorized"}}
	ErrorBanDB    = &RequestDate{Code: http.StatusInternalServerError, Data: &gin.H{"Code": "500", "Message": "db error"}}
	ErrorBadService = &RequestDate{Code: http.StatusInternalServerError, Data: &gin.H{"Code": "500", "Message": "Server Error"}}
)


/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-25
* Time: 上午11:43
* */
package response

import (
	"MsgGO/common/defs"
	"github.com/gin-gonic/gin"
)

func ReqMsg(ctx *gin.Context, data *defs.RequestDate) {
	ctx.JSON(data.Code, data)
}

func ReqInpMsg(ctx *gin.Context, code int, msg string) {
	date := &defs.RequestDate{Code: code, Data: &gin.H{"Message": msg}}
	ReqMsg(ctx, date)
}

func ReqInpData(ctx *gin.Context, code int, data interface{}) {
	date := &defs.RequestDate{Code: code, Data: data}
	ReqMsg(ctx, date)
}


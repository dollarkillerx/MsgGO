/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-25
* Time: 上午11:07
* */
package container

import (
	"MsgGO/common/defs"
	"MsgGO/common/response"
	"github.com/gin-gonic/gin"
)

func RegisterApp(ctx *gin.Context) {
	s, b := ctx.GetPostForm("app_name")
	// 校验
	if b==false {
		response.ReqMsg(ctx,defs.ErrorBadRequest)
		return
	}

	// 数据入库


}

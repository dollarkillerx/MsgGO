/**
* @Author: dollarkiller
* @Date: 19-5-30 下午1:48
* @Description:
* */
package defs

type BaseResult struct {
	HttpSc int
	Data map[string]interface{}
}

var (
	CheckAuthError = &BaseResult{HttpSc:400,Data: map[string]interface{}{"data":"check auth Non-existent","code":"007"}}
	CheckAuthBeOverdue = &BaseResult{HttpSc:400,Data: map[string]interface{}{"data":"you jwt be overdue","code":"008"}}
	CheckAuthOK = &BaseResult{HttpSc:200,Data: map[string]interface{}{"data":"success","code":"200"}}
)
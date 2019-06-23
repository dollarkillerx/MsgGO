/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 上午10:29
* */
package defs

type RequestDate struct {
	Code int
	Data interface{}
}

var (
	ErrorBadRequest = &RequestDate{Code: 400, Data: &SMap{"Code": "400", "Message": "request bad"}}
)

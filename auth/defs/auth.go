/**
* @Author: dollarkiller
* @Date: 19-5-29 上午10:08
* @Description:
* */
package defs

const (
	TokenOverdue = "overdue"
	TokenInvalid = "invalid"
	TokenOk = "ok"
)

// 载荷（Payload）
type Payload struct {
	User string `json:"user"`
	Iat string `json:"iat"` // 签发时间
	Exp string `json:"exp"` // 过期时间
}

// 头部（Header）
type Header struct {
	Type string `json:"type"` // 签发类型
}

package main

type Message struct {
	Id int64 `json:"id,omitempty" form:"id"` //消息id
	Userid int64 `json:"userid,omitempty" form:"userid"`//发送用户
	Cmd int `json:"cmd,omitempty" form:"cmd"`//群聊还是私聊
	Dstid int64 `json:"dstid,omitempty" form:"dstid"` // 对端id or 群id
	Media int `json:"media,omitempty" form:"media"` //消息样式
	Content string `json:"content,omitempty" form:"content"` //消息内容
	Pic string `json:"pic,omitempty" form:"pic"` //预览图片
	Url string `json:"url,omitempty" form:"url"` //服务url
	Memo string `json:"memo,omitempty" form:"memo"` //简单描述
	Amount int `json:"amount,omitempty" form:"amount"` //和数字相关的
}

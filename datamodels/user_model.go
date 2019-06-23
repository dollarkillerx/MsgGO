/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 上午11:43
* */
package datamodels

const (
	SEX_WOMEN  = "W"
	SEX_MEN    = "M"
	SEX_UNKNOW = "U"
)

type User struct {
	//用户ID
	Id       int64  `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Mobile   string `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd   string `xorm:"varchar(40)" form:"passwd" json:"-"`
	Avatar   string `xorm:"varchar(150)" form:"avatar" json:"avatar"` // 头像
	Sex      string `xorm:"varchar(2)" form:"sex" json:"sex"`            // 性别
	Nickname string `xorm:"varchar(20)" form:"nickname" json:"nickname"` // 昵称
	//加盐随机字符串6
	Salt   string `xorm:"varchar(10)" form:"salt" json:"-"`
	Online int    `xorm:"int(10)" form:"online" json:"online"` //是否在线
	//前端鉴权因子,
	Token    string    `xorm:"varchar(40)" form:"token" json:"token"`
	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`     // 群组
	Createat int64 `xorm:"int" form:"createat" json:"createat"` // 创建时间
}

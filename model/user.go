package model

import (
	"time"
)

const (
	SEX_WOMEN = "W" // 女性
	SEX_MAN = "M" // 男性
	SEX_UNKNOW = "U" //未知
)

type User struct {
	Id         int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`//用户id
	Mobile   string 		`xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd       string	`xorm:"varchar(40)" form:"passwd" json:"-"`   // psd + salt
	Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"`//头像
	Sex        string	`xorm:"varchar(2)" form:"sex" json:"sex"`
	Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`
	Salt       string	`xorm:"varchar(10)" form:"salt" json:"-"` // salt
	Online     int	`xorm:"int(10)" form:"online" json:"online"` //是否在线
	Token      string	`xorm:"varchar(40)" form:"token" json:"token"` //鉴权
	Memo      string	`xorm:"varchar(140)" form:"memo" json:"memo"` //备注
	Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"` //统计用户增量
}


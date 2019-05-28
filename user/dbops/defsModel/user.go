package defsModel

type ImUser struct {
	Id int64 `xorm:"pk autoincr"`
	Name string `xorm:"unique"`
	Password string `xorm:"char(32)"`
	Salt string `xorm:"char(32)"`
	Token string `xorm:"char(64)"`
}

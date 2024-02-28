package models

import "time"

type Model struct {
	Id      int64     `xorm:"pk autoincr"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type UserBasic struct {
	Model    `xorm:"extends"`
	UserName string `xorm:"varchar(25) notnull unique 'username' comment('姓名')"`
	Password string `xorm:"varchar(255) notnull unique 'password' comment('密码')"`
	Email    string `xorm:"varchar(25) notnull unique 'email' comment('邮箱')" `
	Name     string `xorm:"varchar(20) notnull default '' 'name'"`
	Status   int8   `xorm:"comment('封禁状态，注册默认为0，后续通过管理员修改') default(0) "`
	IsAdmin  int8   `xorm:"tinyint notnull default 0 'is_admin' comment('权限，后面可以用超级管理员修改') "`
	Avatar   string `xorm:"comment('头像')"`
	Gender   int    `xorm:"comment('性别0为男，1为女）'"`
}

func (table *UserBasic) UserBasic() string {
	return "user"
}

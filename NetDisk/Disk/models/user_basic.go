package models

import "time"

type Model struct {
	Id       int64     `xorm:"pk autoincr"`
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	DeleteAt time.Time `xorm:"deleted" `
}
type UserBasic struct {
	Model    `xorm:"extends"`
	UserName string `xorm:"varchar(25) notnull unique 'username' comment('姓名')"`
	Password string `xorm:"varchar(255) notnull unique 'password' comment('密码')"`
	Email    string `xorm:"varchar(25) notnull unique 'email' comment('邮箱')" `
	Phone    string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Status   int8   `xorm:"comment('权限，注册默认为0，后续通过超级用户可以修改权限') default(0) "`
	Avatar   string `xorm:"comment('头像')"`
	Gender   int    `xorm:"comment('性别0为男，1为女）'"`
}

func (table *UserBasic) UserBasic() string {
	return "user_basic"
}

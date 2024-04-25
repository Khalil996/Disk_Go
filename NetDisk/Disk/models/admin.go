package models

type Admin struct {
	Model    `xorm:"extends"`
	Name     string `xorm:"varchar(20) notnull default '' 'name'"`
	Username string `xorm:"varchar(20) notnull unique 'username' comment('账号')"`
	Password string `xorm:"varchar(255) notnull default '' 'password'"`
	Status   int8   `xorm:"tinyint notnull default 0 'status'"`
}

func (table *Admin) Admin() string {
	return "admin"
}

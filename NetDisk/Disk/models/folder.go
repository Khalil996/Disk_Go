package models

// Folder 网盘文件夹
type Folder struct {
	Model    `xorm:"extends"`
	ParentId int64  `xorm:"bigint notnull default 0 'parent_id' comment('父文件夹id')"`
	Name     string `xorm:"varchar(64) notnull default '' 'name' comment('文件夹名')"`
	UserId   int64  `xorm:"bigint notnull default 0 'user_id'"`
	DelFlag  int8   `xorm:"tinyint notnull default 0 'del_flag' comment('文件删除状态：0：未删除，1：删除')"`
}

func (*Folder) TableName() string {
	return "folder"
}

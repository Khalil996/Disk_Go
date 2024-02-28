package models

import "time"

type File struct {
	Model    `xorm:"extends"`
	UserId   int64     `xorm:"bigint notnull default 0 'user_id'"`
	FsId     int64     `xorm:"bigint notnull default 0 'fs_id'"`
	FolderId int64     `xorm:"bigint notnull default 0 'folder_id'"`
	Name     string    `xorm:"varchar(255) notnull default '' 'name' comment('用户视角文件名')"`
	Url      string    `xorm:"varchar(255) notnull default '' 'url' comment('')"`
	Type     int8      `xorm:"tinyint notnull default 0 comment('类别')"`
	Size     int64     `xorm:"bigint notnull default 0 'size' comment('文件大小')"`
	Status   int8      `xorm:"tinyint notnull default 0 'status' comment('文件状态，0：待合并/未上传，1：上传成功')"`
	IsBig    int8      `xorm:"tinyint notnull default 0 'is_big' comment('是否大文件，0：不是，1：是')"`
	DoneAt   time.Time `xorm:"datetime 'done_at' comment('大文件合并完成时间')"`
	DelFlag  int8      `xorm:"tinyint notnull default 0 'del_flag' comment('文件删除状态：2：未删除，3：删除')"`
}

func (table *File) TableName() string {
	return "file"
}

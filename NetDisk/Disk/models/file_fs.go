package models

import "time"

// FileFs 实际存储
type FileFs struct {
	Model      `xorm:"extends"`
	Bucket     string    `xorm:"varchar(255) notnull default '' bucket comment('桶名')"`
	Ext        string    `xorm:"varchar(255) notnull default '' 'ext' comment('文件扩展名')"`
	ObjectName string    `xorm:"varchar(255) notnull default '' 'object_name' comment('存储路径名')"`
	Hash       string    `xorm:"varchar(255) notnull default '' 'hash' comment('哈希值')"`
	Name       string    `xorm:"varchar(255) notnull default '' 'name' comment('实际文件名')"`
	Size       int64     `xorm:"bigint notnull default 0 'size' comment('文件大小')"`
	ChunkNum   int64     `xorm:"bigint notnull default 0 'chunk_num' comment('分片数量')"`
	Url        string    `xorm:"varchar(255) notnull default '' 'url' comment('访问地址')"`
	Status     int8      `xorm:"tinyint notnull default 0 'status' comment('文件状态，0：大文件未上传，1：大文件待合并，2：小文件未上传，3：上传成功')"` //
	DoneAt     time.Time `xorm:"datetime 'done_at' comment('大文件合并完成时间')"`
}

// TableName 设置表名
func (u *FileFs) TableName() string {
	return "file_fs"
}

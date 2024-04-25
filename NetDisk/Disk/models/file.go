package models

import "time"

type File struct {
	Model      `xorm:"extends"`
	UserId     int64     `xorm:"bigint notnull default 0 'user_id'" json:"userId"`
	FsId       int64     `xorm:"bigint notnull default 0 'fs_id'" json:"fsId"`
	FolderId   int64     `xorm:"bigint notnull default 0 'folder_id'" json:"folderId"`
	Name       string    `xorm:"varchar(255) notnull default '' 'name' comment('用户视角文件名')" json:"name"`
	ObjectName string    `xorm:"varchar(255) notnull default '' 'object_name' comment('存储路径名')" json:"objectName"`
	Ext        string    `xorm:"varchar(255) notnull default '' 'ext' comment('文件扩展名')" json:"ext"`
	Size       int64     `xorm:"bigint notnull default 0 'size' comment('文件大小')" json:"size"`
	Type       int8      `xorm:"tinyint notnull default 0 comment('类别')" json:"type"`
	Status     int8      `xorm:"tinyint notnull default 0 'status' comment('文件状态，0：待合并/未上传，1：上传成功')" json:"status"`
	IsBig      int8      `xorm:"tinyint notnull default 0 'is_big' comment('是否大文件，0：不是，1：是')" json:"isBig"`
	DoneAt     time.Time `xorm:"datetime 'done_at' comment('大文件合并完成时间')" json:"doneAt"`
	DelFlag    int8      `xorm:"tinyint notnull default 0 'del_flag' comment('文件删除状态：2：未删除，3：删除')" json:"delFlag"`
	DelTime    int64     `xorm:"bigint notnull default 0 'del_time' comment('')" json:"delTime"`
	SyncFlag   int8      `xorm:"tinyint notnull default 0 'sync_flag' comment('同步es标志')" json:"syncFlag"`
}

func (table *File) TableName() string {
	return "file"
}

package models

type ShareFile struct {
	Model   `xorm:"extends"`
	ShareId string `xorm:"varchar(255) notnull default '' 'share_id' comment('分享id')"`
	FileId  int64  `xorm:"bigint notnull default 0 'file_id'"`
}

func (table *ShareFile) ShareFile() string {
	return "share_file"
}

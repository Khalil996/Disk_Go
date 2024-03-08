package models

type FileShare struct {
	Model       `xorm:"extends"`
	UserId      int64 `xorm:"bigint notnull default 0 'user_id'"`
	FileId      int64 `xorm:"bigint notnull default 0 'file_id'"`
	Expired     int64 `xorm:"bigint notnull default 0 'expired'"`
	DownloadNum int64 `xorm:"bigint notnull default 0 'download_num'"`
	ClickNum    int64 `xorm:"bigint notnull default 0 'click_num'"`
	Status      int8  `xorm:"tinyint notnull default 0 'status'"`
}

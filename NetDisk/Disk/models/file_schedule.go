package models

// FileSchedule 任务表
type FileSchedule struct {
	Model    `xorm:"extends"`
	FileId   int64 `xorm:"bigint notnull default 0 'file_id'"`
	FsId     int64 `xorm:"bigint notnull default 0 'fs_id'"`
	ChunkNum int64 `xorm:"bigint notnull default 0 'chunk_num'"`
}

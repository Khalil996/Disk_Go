package redis

import "time"

const (
	NetdiskCapacity = "capacity"

	UserLogin      = "user:login:"
	EmailValidCode = "user:code:"
	UserInfoKey    = "user:info:"

	UploadCheckKey        = "upload:file:"
	UploadCheckBigFileKey = "upload:file-shard:"
	UploadCheckChunkKeyF  = "upload:file-shard:%d:chunk:%d"

	FileFolderDownloadUrlKey = "download:folder:%d:%d" // download:folder:{userId}:{folderId}
	FileTypeDownloadUrlKey   = "download:type:%d:%d"   // download:type:{userId}:{type}
)

const (
	RegisterCodeExpire = 5*time.Minute + 10*time.Minute
	UserInfoExpire     = 7 * 24 * time.Hour

	UploadCheckExpire      = 24 * time.Hour
	UploadCheckChunkExpire = 10 * time.Minute

	DownloadExpire = 7*24*time.Hour - 10000
)

package redis

import "time"

const (
	UserLogin = "user:login:"

	UploadCheckKey        = "upload:file:"
	UploadCheckBigFileKey = "upload:file-shard:"
	UploadCheckChunkKeyF  = "upload:file-shard:%d:chunk:%d"

	FileFolderDownloadUrlKey = "download:folder:%d:%d" // download:folder:{userId}:{folderId}
	FileTypeDownloadUrlKey   = "download:type:%d:%d"   // download:type:{userId}:{type}
)

const (
	UploadCheckExpire      = 24 * time.Hour
	UploadCheckChunkExpire = 10 * time.Minute
)

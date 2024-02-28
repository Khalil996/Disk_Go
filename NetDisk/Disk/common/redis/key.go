package redis

import "time"

const (
	UserLogin = "user:login:"

	UploadCheckKey        = "upload:file:"
	UploadCheckBigFileKey = "upload:file-shard:"
	UploadCheckChunkKeyF  = "upload:file-shard:%d:chunk:%d"

	DownloadGetFsKey = "download:fs:fn:"
)

const (
	UploadCheckExpire      = 24 * time.Hour
	UploadCheckChunkExpire = 10 * time.Minute
)

package define

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	Id   int64
	Name string
	jwt.RegisteredClaims
}

var JwtKey = "Disk-key"

// token过期时间
var TokenExpire = 7200

// token刷新过期时间
var RefreshTokenExpire = 7200

// 验证码长度
var CodeLength = 6

// 验证码过期时间
var CodeExpire = 300

const (
	UserIdKey   = "userId"
	UserNameKey = "userName"
)

const (
	ShardingFloor float64 = 20971520 // 需要分片起始大小：20MB
	ShardingSize  float64 = 5242880  // 分片大小： 5MB
)

const (
	SmallFileFlag int8 = iota
	BigFileFlag
)

// 0：待合并/未上传，1：上传成功
const (
	StatusFileUnuploaded int8 = iota
	StatusFileUploaded
	StatusFileUndeleted
	StatusFileDeleted
)

const (
	ConfirmNotShard = iota
	ConfirmShard
)

// 0：大文件未上传，1：大文件待合并，2：小文件未上传，3：上传成功
const (
	StatusFsBigFileUnuploaded int8 = iota
	StatusFsBigFileNeedMerge
	StatusFsFileUnuploaded
	StatusFsUploaded
)

const (
	StatusFolderUndeleted int8 = iota
	StatusFolderDeleted
)

// context keys
const (
	CtxFolderIdsKey = "folderIds"
	CtxFileIdsKey   = "fileIds"
)

const (
	TimeFormat1 = "2006-01-02 15:04:05"
	TimeFormat2 = "2006-01-02/15:04:05"
)

const (
	TypeDocs = iota
	TypeImage
	TypeVideo
	TypeAudio
	TypeOther
)

package define

import (
	"regexp"
)

var (
	Upattern, _ = regexp.Compile("^[a-zA-Z0-9]{4,20}$")
	Ppattern, _ = regexp.Compile("^[a-zA-Z0-9!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~]{6,20}$")

	BanM = map[int8]string{
		StatusUserBannedByName:      "昵称",
		StatusUserBannedByUsername:  "用户名",
		StatusUserBannedBySignature: "签名",
		StatusUserBannedByAvatar:    "头像",
		StatusUserBannedByShare:     "分享",
	}
)

var (
	DocSuffix = map[string]bool{
		".doc":  true,
		".docx": true,
		".xlsx": true,
		".xls":  true,
		".ppt":  true,
		".pptx": true,
		".pdf":  true,
		".txt":  true,
		".md":   true,
	}

	ImageSuffix = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".tif":  true,
		".tiff": true,
		".webp": true,
	}

	AudioSuffix = map[string]bool{
		".mp3":  true,
		".wav":  true,
		".aac":  true,
		".flac": true,
		".ogg":  true,
		".wma":  true,
		".m4a":  true,
		".amr":  true,
		".ape":  true,
	}

	VideoSuffix = map[string]bool{
		".avi":  true,
		".rm":   true,
		".rmvb": true,
		".mpeg": true,
		".mpg":  true,
		".mp4":  true,
		".wmv":  true,
		".mov":  true,
		".flv":  true,
		".mkv":  true,
		".3gp":  true,
		".m4v":  true,
	}
)

func GetTypeByBruteForce(ext string) int8 {
	if DocSuffix[ext] {
		return TypeDocs
	} else if ImageSuffix[ext] {
		return TypeImage
	} else if AudioSuffix[ext] {
		return TypeAudio
	} else if VideoSuffix[ext] {
		return TypeVideo
	} else {
		return TypeOther
	}
}

var ShareExpireType = map[int8]int64{
	ShareExpireDay:     86400,
	ShareExpireWeek:    604800,
	ShareExpireMonth:   2592000,
	ShareExpireForever: 0,
}

var OperationM = map[string]string{
	"Login":               "账号：%v 登录了",
	"Upload":              "上传了文件，文件id：%v",
	"CopyFiles":           "复制文件id：%v",
	"CopyFolders":         "复制了文件夹，文件夹id：%v",
	"CreateFolder":        "创建了文件夹，名称：%v",
	"DeleteAllFilesTruly": "清空了回收站",
	"DeleteFiles":         "删除了文件，文件id：%v",
	"DeleteFilesTruly":    "删除了回收站文件，文件id：%v",
	"DeleteFolders":       "删除文件夹，文件夹id：%v",
	"MoveFiles":           "移动了文件，文件id：%v",
	"MoveFolders":         "移动了文件夹，文件夹id：%v",
	"RecoverFiles":        "恢复了文件id：%v",
	"ShareCancel":         "取消了分享，分享id：%v",
	"ShareFolder":         "分享了文件夹，分享的文件夹id：%v",
	"Share":               "分享了文件，分享的文件id：%v",
	"UpdateFile":          "修改了文件名，文件id：%v",
	"UpdateFolder":        "修改了文件夹名，文件夹id：%v",
	"AdminLogin":          "管理员账号：%v 登录了",
	"SetFileStatus":       "管理员设置文件状态：%v",
	"SetShareStatus":      "管理员设置分享状态：%v",
	"SetUserStatus":       "设置用户状态：%v",
	"SetAdminStatus":      "设置管理员状态：%v",
	"MergeLogic":          "分片文件合并任务开始，时间：%v",
	"Search":              "搜索文件，搜索词条：%v",
	"UpdateDetail":        "更新信息：%v",
	"UpdateAvatar":        "更新头像：%v",
	"AddAdmin":            "新增了管理员：%v",
	"DeleteAdmin":         "删除了管理员：%v",
}

package common

import (
	"cloud_go/Disk/internal/types"
	"net/http"
)

// 获取单个文件
func GetSingleFile(r *http.Request) (*types.FileParam, error) {
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}

	return &types.FileParam{
		File:       file,
		FileHeader: header,
	}, nil
}

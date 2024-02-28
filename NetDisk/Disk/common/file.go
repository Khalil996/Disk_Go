package common

import (
	"cloud_go/Disk/internal/types"
	"net/http"
)

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

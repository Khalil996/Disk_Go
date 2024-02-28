package types

import "mime/multipart"

type FileParam struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

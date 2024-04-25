package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/url"
	"os"
	"strings"
	"time"
)

type (
	Service struct {
		BucketName string
		client     *minio.Client
	}
)

func (c *Client) NewService() *Service {
	return &Service{c.BucketName, c.client}
}

// Upload 上传文件
func (s *Service) Upload(ctx context.Context, objectName string, file io.Reader) error {
	_, err := s.client.PutObjectWithContext(ctx, s.BucketName, objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		logx.Errorf("minio-上传文件出错，err: %v", err)
		return err
	}
	return nil
}

// FUpload 上传文件 by file
func (s *Service) FUpload(ctx context.Context, objectName, filePath string) error {
	_, err := s.client.FPutObjectWithContext(ctx, s.BucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		logx.Errorf("minio-上传文件出错，err: %v", err)
		return err
	}
	return nil
}

func (s *Service) IfExist(objectName string) (bool, int64, error) {
	info, err := s.client.StatObject(s.BucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "The specified key does not exist") {
			return false, 0, nil
		}
		logx.Errorf("minio-判断文件是否存在出错，err: %v", err)
		return false, 0, err
	}
	return true, info.Size, nil
}

func (s *Service) GenUrl(objectName string, filename string, download bool) (string, error) {
	var (
		u   *url.URL
		err error
	)

	if download {
		kvs := url.Values{}
		kvs["response-content-disposition"] = []string{"attachment; filename=" + objectName}
		if filename != "" {
			kvs["response-content-disposition"] = []string{"attachment; filename=" + filename}
		}
		u, err = s.client.PresignedGetObject(s.BucketName, objectName, 7*24*time.Hour, kvs)
	} else {
		u, err = s.client.PresignedGetObject(s.BucketName, objectName, 7*24*time.Hour, nil)
	}

	if err != nil {
		logx.Errorf("minio-获取下载url出错，err: %v", err)
		return "", err
	}
	return fmt.Sprintf("%v", u), nil
}

// DownloadChunk2Local 下载文件切片
func (s *Service) DownloadChunk2Local(ctx context.Context, objectName, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		logx.Errorf("DownloadChunk2Local，创建临时文件：[%v]，出错，ERR: [%v]", err)
		return err
	}
	file.Close()

	if err = s.client.FGetObjectWithContext(ctx, s.BucketName, objectName,
		filename, minio.GetObjectOptions{}); err != nil {
		logx.Errorf("DownloadChunk2Local，minio下载出错，ERR: [%v]", err)
		return err
	}
	return nil
}

// DeleteFile 删除文件
func (s *Service) DeleteFile(objectName string) error {
	err := s.client.RemoveObject(s.BucketName, objectName)
	if err != nil {
		logx.Errorf("DeleteFile，删除分片出错，ERR: [%v]", err)
		return err
	}
	return err
}

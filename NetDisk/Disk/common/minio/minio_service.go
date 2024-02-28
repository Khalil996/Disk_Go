package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/url"
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
	fmt.Println(s.BucketName)
	if err != nil {
		logx.Errorf("minio-上传文件出错，err: %v", err)
		return err
	}
	return nil
}

func (s *Service) IfExist(objectName string) (bool, error) {
	_, err := s.client.StatObject(s.BucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "The specified key does not exist") {
			return false, nil
		}
		logx.Errorf("minio-判断文件是否存在出错，err: %v", err)
		return false, err
	}
	return true, nil
}

func (s *Service) GenUrl(objectName string, download bool) (string, error) {
	var u *url.URL
	var err error

	if download {
		kvs := url.Values{"response-content-disposition": []string{"attachment; filename=" + objectName}}
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

//// DownloadFile 下载文件
//func (s *Service) DownloadFile(ctx context.Context, objectName string) (string, error) {
//
//	file, err := os.CreateTemp("/tmp/netdisk/", "*")
//	if err != nil {
//		return "", err
//	}
//	defer file.Close()
//
//	filename := file.Name()
//	if err = s.client.FGetObjectWithContext(ctx, s.BucketName, objectName,
//		filename, minio.GetObjectOptions{}); err != nil {
//		logx.Errorf("minio下载url出错，err: %v", err)
//		return "", err
//	}
//
//	return filename, nil
//}
//
//// DeleteFile 删除文件
//func (s *Service) DeleteFile(bucketName, objectName string) (bool, error) {的miniokehuduan1
//	// 删除存储桶中的文件
//	err := s.Client.RemoveObject(bucketName, objectName)
//	if err != nil {
//		log.Println("remove object fail: ", err)
//		return false, err
//	}
//
//	fmt.Println("Successfully deleted", objectName)
//	return true, err
//}
//
//// ListObjects 列出文件
//func (s *Service) ListObjects(bucketName, prefix string) ([]string, error) {
//	var objectNames []string
//
//	for object := range s.Client.ListObjects(bucketName, prefix, true, nil) {
//		if object.Err != nil {
//			return nil, object.Err
//		}
//
//		objectNames = append(objectNames, object.Key)
//	}
//
//	return objectNames, nil
//}

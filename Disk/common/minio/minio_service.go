package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"log"
	"os"
	"strings"
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
		log.Println("putObject fail: ", err)
		return err
	}
	return nil
}

func (s *Service) IfExist(objectName string) (bool, error) {
	_, err := s.client.StatObject(s.BucketName, objectName, minio.StatObjectOptions{})
	if strings.Contains(err.Error(), "The specified key does not exist") {
		return false, nil
	} else if err != nil {
		log.Println("statObject fail: ", err)
		return false, err
	}
	return true, nil
}

// DownloadFile 下载文件
func (s *Service) DownloadFile(ctx context.Context, objectName string) (string, error) {

	file, err := os.CreateTemp("/tmp/netdisk/", "*")
	if err != nil {
		return "", err
	}
	defer file.Close()

	filename := file.Name()
	if err = s.client.FGetObjectWithContext(ctx, s.BucketName, objectName,
		filename, minio.GetObjectOptions{}); err != nil {
		return "", err
	}

	return filename, nil
}

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

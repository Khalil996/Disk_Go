package minio

import (
	"github.com/minio/minio-go"
	"log"
	"strconv"
	"time"
)

type (
	Conf struct {
		MinioPort      string
		MinioAccessKey string
		MinioSecretKey string
		MinioBucket    string
		MinioUseSSL    bool
	}

	Client struct {
		BucketName string
		client     *minio.Client
	}
)

// Init 初使化 minio client对象。
func Init(conf *Conf) *Client {
	client, err := minio.New(conf.MinioPort, conf.MinioAccessKey, conf.MinioSecretKey, conf.MinioUseSSL)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	bucketName := conf.MinioBucket
	exists, err := client.BucketExists(bucketName)
	if err != nil || !exists {
		err = client.MakeBucket(bucketName, "")
		if err != nil {
			log.Fatalln(err)
		}
	}

	return &Client{bucketName, client}
}

// GenObjectName 生成objectName
func (c *Client) GenObjectName(hash string, ext string) (string, string) {
	filename := hash + ext
	return filename, "/" + time.Now().Format("2006-01") + "/" +
		string(hash[0]) + "/" + string(hash[0]) + "/" + hash + "/" + filename
}

func (c *Client) GenChunkObjectName(hash string, chunkId int64) string {
	return "/" + time.Now().Format("2006-01") + "/" +
		string(hash[0]) + "/" + string(hash[0]) + "/" + hash +
		"/" + hash + "@" + strconv.FormatInt(chunkId, 10)
}

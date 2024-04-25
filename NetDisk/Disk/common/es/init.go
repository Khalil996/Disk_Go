package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

var Es *elastic.Client

func Init(esHost string) *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetURL(esHost),
		elastic.SetSniff(false))

	if err != nil {
		log.Fatalf("初始化 ES 失败，ERR: [%v]\n", err)
		return nil
	}

	info, code, err := client.Ping(esHost).Do(context.Background())
	if err != nil {
		log.Fatalf("PING ES 失败，ERR: [%v]\n", err)
	}
	logx.Info(info, " ", code)

	Es = client
	return client
}

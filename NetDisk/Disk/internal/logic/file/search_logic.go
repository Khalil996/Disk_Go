package file

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/logic/mqs"
	"cloud_go/Disk/models"
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"

	"cloud_go/Disk/internal/svc"
	"cloud_go/Disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (interface{}, error) {
	// todo: add your logic here and delete this line
	var (
		userId = l.ctx.Value(define.UserIdKey).(int64)
		err    error
	)
	defer mqs.LogSend(l.ctx, err, "search", req.Phrase)

	bq := elastic.NewBoolQuery()

	bq1 := elastic.NewBoolQuery().Should(
		elastic.NewMatchQuery("Name", req.Phrase),
		elastic.NewMatchQuery("Ext", req.Phrase),
	)
	bq = bq.Must(
		elastic.NewTermQuery("UserId", userId),
		bq1,
	)
	do, err := l.svcCtx.Es.Search().Index("netdisk_file").Query(bq).Do(context.TODO())
	if err != nil {
		logx.Errorf("Elasticsearch query failed: %v", err)
		return nil, err
	}

	if do == nil || do.Hits == nil {
		return nil, errors.New("received nil response or hits from Elasticsearch")
	}

	var files []*models.File
	if len(do.Hits.Hits) > 0 {
		for _, hit := range do.Hits.Hits {
			var file models.File
			err = json.Unmarshal(hit.Source, &file)
			if err != nil {
				logx.Error("Search,file:[%v] 反序列化失败,Err:[%v]", file.Id, err)
				continue
			}
			files = append(files, &file)
		}
	} else {
		if err = l.svcCtx.Engine.Where("user_id = ?", userId).
			And("name like ?", "%"+req.Phrase+"%").
			Find(&files); err != nil {
			err = errors.New("Search，查询数据库失败，ERR: " + err.Error())
			logx.Errorf(err.Error())
		}
	}

	return files, nil
}

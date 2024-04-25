package common

import (
	"cloud_go/Disk/common/es"
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

var fileIndex = "netdisk_file"

func SyncLogic() {
	logx.Infof("同步任务开始")

	var files []*models.File
	if err := xorm.Xorm.
		Where("sync_flag > ?", define.FlagSyncDone).
		Limit(1000).Find(&files); err != nil {
		logx.Errorf("SyncLogic，查询file出错，ERR：[%v]", err)
		return
	}

	q := make(chan struct{}, 5)
	for _, file := range files {
		f := file
		go func() {
			q <- struct{}{}
			Sync(f, func(int64) {})
			<-q
		}()
	}
}

func Sync(file *models.File, errCallBack func(int64)) {
	var (
		ctx = context.Background()
		err error
	)

	defer func() {
		if err != nil {
			errCallBack(file.Id)
		}
	}()

	switch file.SyncFlag {
	case define.FlagSyncWrite:
		if _, err = xorm.Xorm.DoTransaction(writeSync(ctx, file)); err != nil {
			logx.Errorf("SyncLogic，同步更新失败，ERR: [%v]", err)
		}

	case define.FlagSyncDelete:
		if _, err = xorm.Xorm.DoTransaction(deleteSync(ctx, file)); err != nil {
			logx.Errorf("SyncLogic，同步删除失败，ERR: [%v]", err)
		}
	}

}

func writeSync(ctx context.Context, file *models.File) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {
		if _, err := updateDb(session, file.Id); err != nil {
			return nil, err
		}

		if _, err := es.Es.Index().
			Index(fileIndex).
			Id(strconv.FormatInt(file.Id, 10)).
			BodyJson(file).
			Do(ctx); err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func deleteSync(ctx context.Context, file *models.File) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {
		if _, err := updateDb(session, file.Id); err != nil {
			return nil, err
		}

		if _, err := es.Es.Delete().
			Index(fileIndex).
			Id(strconv.FormatInt(file.Id, 10)).Do(ctx); err != nil {
		}
		return nil, nil
	}
}

func updateDb(session *xorm.Session, id int64) (int64, error) {
	file := &models.File{SyncFlag: define.FlagSyncDone}
	return session.ID(id).Cols("sync_flag").Update(file)
}

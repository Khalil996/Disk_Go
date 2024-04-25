package cron

import (
	"cloud_go/Disk/common"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"time"
)

func SyncTask() {

	timezone, _ := time.LoadLocation("Asia/Shanghai")
	mergeCron := cron.New(cron.WithLocation(timezone))

	_, err := mergeCron.AddFunc("*/3 * * * *", common.SyncLogic)
	if err != nil {
		log.Fatalf("SyncTask，添加同步定时任务失败，ERR: [%v]", err)
	}

	logx.Info("SyncTask，添加同步定时任务成功")
	mergeCron.Start()
}

package mqs

import (
	"cloud_go/Disk/define"
	"cloud_go/Disk/internal/config"
	"context"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	errKqPusher       = "kq push error"
	operationKqPusher = "operation"
)

var m = make(map[string]*kq.Pusher)

func NewLogPusher(confs []*config.KqPusherConf) {
	for _, conf := range confs {
		m[conf.Type] = kq.NewPusher(conf.Brokers, conf.Topic)
	}
}

func LogSend(ctx context.Context, err error, operation string, ids ...interface{}) {
	ErrSend(operation, err)
	OperationSend(ctx, operation, err, ids)
}

func ErrSend(operation string, err error) {
	if err != nil {
		msg := err.Error()
		if err2 := m[errKqPusher].Push(operation + msg); err2 != nil {
			logx.Errorf("Kafka发送错误日志失败, Msg: [%v], ERR: [%v]", msg, err2)
		}
	}
}

func OperationSend(ctx context.Context, operation string, err error, ids ...interface{}) {
	var (
		userId  int64
		name    string
		content string
		op      = define.OperationM[operation]
		prefix  = "操作成功，"
	)

	if err != nil {
		prefix = "操作失败，"
	}

	if ids != nil {
		op = fmt.Sprintf(define.OperationM[operation], ids)
	}

	if v := ctx.Value(define.UserIdKey); v == nil {
		content = op
		logx.Error(2, v, content, op)
	} else {
		userId = ctx.Value(define.UserIdKey).(int64)
		name = ctx.Value(define.UserNameKey).(string)
		logx.Error(1, v, content, op)
		content = fmt.Sprintf(define.Operation, userId, name, op)
	}

	if err = m[operationKqPusher].Push(prefix + content); err != nil {
		logx.Errorf("Kafka发送操作日志失败, Msg: [%v], ERR: [%v]", content, err)
	}
}

package xorm

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Xorm *Engine

type (
	DbConf struct {
		Dsn string
	}

	Engine struct {
		*xorm.Engine
	}

	Session struct {
		*xorm.Session
	}

	TxFn func(session *Session) (interface{}, error)
)

func InitEngine(conf *DbConf) *Engine {
	engine, err := xorm.NewEngine("mysql", conf.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	if err = engine.Ping(); err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败" + err.Error())
	}

	engine.Logger().ShowSQL(true)
	e := &Engine{engine}
	Xorm = e
	return e
}

func (e *Engine) DoTransaction(fn TxFn) (interface{}, error) {
	return e.Transaction(func(session *xorm.Session) (interface{}, error) {
		return fn(&Session{session})
	})
}

func (e *Engine) DoTransactions(session *Session,
	fns ...func(*Session) (interface{}, error)) ([]interface{}, error) {
	var results []interface{}
	if session == nil {
		session = &Session{e.NewSession()}
	}

	if err := session.Begin(); err != nil {
		return nil, err
	}

	defer func() {
		var err error
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
			// TODO: log
		}

		if err != nil {
			_ = session.Rollback()
		} else {
			_ = session.Commit()
		}
	}()

	for _, fn := range fns {
		res, err := fn(session)
		if err != nil {
			return nil, err
		}
		if res != nil {
			results = append(results, res)
		}
	}

	return results, nil
}

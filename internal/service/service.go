package service

import (
	"context"

	"github.com/EZ4BRUCE/rule-server/global"
	"github.com/EZ4BRUCE/rule-server/internal/dao"
)

// Service封装上下文以及dao实例用于执行dao方法
type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}

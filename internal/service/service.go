package service

import (
	"context"

	"github.com/EZ4BRUCE/rule-server/global"
	"github.com/EZ4BRUCE/rule-server/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}

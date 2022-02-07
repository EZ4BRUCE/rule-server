package service

import (
	"github.com/EZ4BRUCE/rule-server/internal/model"
	"github.com/EZ4BRUCE/rule-server/internal/request"
)

// service层方法，接收请求结构体或特定参数执行dao方法

func (svc *Service) CreateRule(param *request.CreateRuleRequest) error {
	return svc.dao.CreateRule(param.Level, param.Action, param.Description)
}

func (svc *Service) DeleteRule(id uint32) error {
	return svc.dao.DeleteRule(id)
}

func (svc *Service) GetRule(id uint32) (model.Rule, error) {
	return svc.dao.GetRule(id)

}

func (svc *Service) UpdateRule(param *request.UpdateRuleRequest) error {
	return svc.dao.UpdateRule(param.Id, param.Level, param.Action, param.Description)
}

func (svc *Service) ListRules() ([]model.Rule, error) {
	return svc.dao.ListRules()
}

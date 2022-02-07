package service

import (
	"github.com/EZ4BRUCE/rule-server/internal/model"
	"github.com/EZ4BRUCE/rule-server/internal/request"
)

// service层方法，接收请求结构体或特定参数执行dao方法

func (svc *Service) CreateAggregator(param *request.CreateAggregatorRequest) error {
	return svc.dao.CreateAggregator(param.Name, param.Metric, param.FunctionId, param.RuleId)
}

func (svc *Service) DeleteAggregator(id uint32) error {
	return svc.dao.DeleteAggregator(id)
}

func (svc *Service) GetAggregator(id uint32) (model.Aggregator, error) {
	return svc.dao.GetAggregator(id)
}

func (svc *Service) UpdateAggregator(param *request.UpdateAggregatorRequest) error {
	return svc.dao.UpdateAggregator(param.Id, param.Name, param.Metric, param.FunctionId, param.RuleId)
}

func (svc *Service) SearchAggregators(metric string) ([]model.Aggregator, error) {
	return svc.dao.SearchAggregators(metric)
}

func (svc *Service) ListAggregators() ([]model.Aggregator, error) {
	return svc.dao.ListAggregators()
}

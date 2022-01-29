package service

import (
	"github.com/EZ4BRUCE/rule-server/internal/model"
	"github.com/EZ4BRUCE/rule-server/internal/request"
)

// Service的方法，svc作为接收者
func (svc *Service) CreateFunction(param *request.CreateFunctionRequest) error {
	return svc.dao.CreateFunction(param.Type, param.Threshold, param.Description)
}

func (svc *Service) DeleteFunction(id uint32) error {
	return svc.dao.DeleteFunction(id)
}

func (svc *Service) GetFunction(id uint32) (model.Function, error) {
	return svc.dao.GetFunction(id)

}

func (svc *Service) UpdateFunction(param *request.UpdateFunctionRequest) error {
	return svc.dao.UpdateFunction(param.Id, param.Type, param.Threshold, param.Description)
}
func (svc *Service) ListFunction() ([]model.Function, error) {
	return svc.dao.ListFunction()
}

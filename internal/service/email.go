package service

import (
	"github.com/EZ4BRUCE/rule-server/internal/model"
	"github.com/EZ4BRUCE/rule-server/internal/request"
)

// service层方法，接收请求结构体或特定参数执行dao方法

func (svc *Service) CreateEmail(param *request.CreateEmailRequest) error {
	return svc.dao.CreateEmail(param.Address)
}
func (svc *Service) DeleteEmail(id uint32) error {
	return svc.dao.DeleteEmail(id)
}

func (svc *Service) GetEmail(id uint32) (model.Email, error) {
	return svc.dao.GetEmail(id)

}

func (svc *Service) UpdateEmail(param *request.UpdateEmailRequest) error {
	return svc.dao.UpdateEmail(param.Id, param.Address)
}

func (svc *Service) ListEmails() ([]model.Email, error) {
	return svc.dao.ListEmails()
}

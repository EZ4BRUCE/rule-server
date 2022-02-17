package service

import (
	"github.com/EZ4BRUCE/rule-server/internal/model"
	"github.com/EZ4BRUCE/rule-server/internal/request"
)

// service层方法，接收请求结构体或特定参数执行dao方法

func (svc *Service) CreateSmtp(param *request.CreateSmtpRequest) error {
	return svc.dao.CreateSmtp(param.Host, param.Port, param.IsSSL, param.UserName, param.Password, param.From)
}

func (svc *Service) DeleteSmtp(id uint32) error {
	return svc.dao.DeleteSmtp(id)
}

func (svc *Service) GetSmtp(id uint32) (model.Smtp, error) {
	return svc.dao.GetSmtp(id)

}

func (svc *Service) UpdateSmtp(param *request.UpdateSmtpRequest) error {
	return svc.dao.UpdateSmtp(param.Id, param.Host, param.Port, param.IsSSL, param.UserName, param.Password, param.From)
}

func (svc *Service) ListSmtps() ([]model.Smtp, error) {
	return svc.dao.ListSmtps()
}

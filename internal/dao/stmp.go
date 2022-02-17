package dao

import "github.com/EZ4BRUCE/rule-server/internal/model"

// dao层方法，接收特定参数执行model方法

func (d Dao) CreateSmtp(host string, port int, isSSL bool, userName, password, from string) error {
	rule := model.Smtp{
		Host:     host,
		Port:     port,
		IsSSL:    isSSL,
		UserName: userName,
		Password: password,
		From:     from,
	}
	return rule.Create(d.engine)
}

func (d Dao) DeleteSmtp(id uint32) error {
	rule := model.Smtp{Id: id}
	return rule.Delete(d.engine)
}

func (d Dao) GetSmtp(id uint32) (model.Smtp, error) {
	rule := model.Smtp{Id: id}
	return rule.Get(d.engine)
}

func (d Dao) UpdateSmtp(id uint32, host string, port int, isSSL bool, userName, password, from string) error {
	rule := model.Smtp{
		Id:       id,
		Host:     host,
		Port:     port,
		IsSSL:    isSSL,
		UserName: userName,
		Password: password,
		From:     from,
	}
	return rule.Update(d.engine)
}

func (d Dao) ListSmtps() ([]model.Smtp, error) {
	rule := model.Smtp{}
	return rule.List(d.engine)
}

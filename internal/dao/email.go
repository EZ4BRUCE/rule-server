package dao

import "github.com/EZ4BRUCE/rule-server/internal/model"

// dao层方法，接收特定参数执行model方法

func (d Dao) CreateEmail(address string) error {
	email := model.Email{Address: address}
	return email.Create(d.engine)
}

func (d Dao) DeleteEmail(id uint32) error {
	email := model.Email{Id: id}
	return email.Delete(d.engine)
}

func (d Dao) GetEmail(id uint32) (model.Email, error) {
	email := model.Email{Id: id}
	return email.Get(d.engine)
}

func (d Dao) UpdateEmail(id uint32, address string) error {
	email := model.Email{Id: id, Address: address}
	return email.Update(d.engine)
}

func (d Dao) ListEmails() ([]model.Email, error) {
	email := model.Email{}
	return email.List(d.engine)
}

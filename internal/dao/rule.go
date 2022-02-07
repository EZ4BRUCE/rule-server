package dao

import "github.com/EZ4BRUCE/rule-server/internal/model"

// dao层方法，接收特定参数执行model方法

func (d Dao) CreateRule(level, action, description string) error {
	rule := model.Rule{Level: level, Action: action, Description: description}
	return rule.Create(d.engine)
}

func (d Dao) DeleteRule(id uint32) error {
	rule := model.Rule{Id: id}
	return rule.Delete(d.engine)
}

func (d Dao) GetRule(id uint32) (model.Rule, error) {
	rule := model.Rule{Id: id}
	return rule.Get(d.engine)
}

func (d Dao) UpdateRule(id uint32, level, action, description string) error {
	rule := model.Rule{Id: id, Level: level, Action: action, Description: description}
	return rule.Update(d.engine)
}

func (d Dao) ListRules() ([]model.Rule, error) {
	rule := model.Rule{}
	return rule.List(d.engine)
}

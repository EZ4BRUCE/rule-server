package dao

import "github.com/EZ4BRUCE/rule-server/internal/model"

// dao的方法，dao作为接收者
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

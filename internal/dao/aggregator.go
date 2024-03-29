package dao

import "github.com/EZ4BRUCE/rule-server/internal/model"

// dao层方法，接收特定参数执行model方法

func (d Dao) CreateAggregator(name, metric string, function_id, rule_id uint32) error {
	aggregator := model.Aggregator{Name: name, Metric: metric, FunctionId: function_id, RuleId: rule_id}
	return aggregator.Create(d.engine)
}

func (d Dao) DeleteAggregator(id uint32) error {
	aggregator := model.Aggregator{Id: id}
	return aggregator.Delete(d.engine)
}

func (d Dao) GetAggregator(id uint32) (model.Aggregator, error) {
	aggregator := model.Aggregator{Id: id}
	return aggregator.Get(d.engine)
}

func (d Dao) UpdateAggregator(id uint32, name, metric string, function_id, rule_id uint32) error {
	aggregator := model.Aggregator{Id: id, Name: name, Metric: metric, FunctionId: function_id, RuleId: rule_id}
	return aggregator.Update(d.engine)
}

func (d Dao) SearchAggregators(metric string) ([]model.Aggregator, error) {
	aggregator := model.Aggregator{}
	return aggregator.Search(d.engine, metric)
}

func (d Dao) ListAggregators() ([]model.Aggregator, error) {
	aggregator := model.Aggregator{}
	return aggregator.List(d.engine)
}

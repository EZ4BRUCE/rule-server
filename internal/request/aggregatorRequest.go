package request

type CreateAggregatorRequest struct {
	Name       string `gorm:"column:name" form:"name"`
	Metric     string `gorm:"column:metric" form:"metric"`
	FunctionId uint32 `gorm:"column:function_id" form:"function_id"`
	RuleId     uint32 `gorm:"column:rule_id" form:"rule_id"`
}

// 其他的请求只有一个参数（ID）或者没有参数（list）故不用构建请求体

type UpdateAggregatorRequest struct {
	Id         uint32 `gorm:"primaryKey;column:id" form:"id"`
	Name       string `gorm:"column:name" form:"name"`
	Metric     string `gorm:"column:metric" form:"metric"`
	FunctionId uint32 `gorm:"column:function_id" form:"function_id"`
	RuleId     uint32 `gorm:"column:rule_id" form:"rule_id"`
}

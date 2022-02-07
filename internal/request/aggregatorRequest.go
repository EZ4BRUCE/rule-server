package request

// 接收聚合器的POST请求参数定义的结构体
// 其他的请求只有一个path参数（ID）或者没有参数（list）故不用构建请求结构体

type CreateAggregatorRequest struct {
	Name       string `gorm:"column:name" form:"name"`
	Metric     string `gorm:"column:metric" form:"metric"`
	FunctionId uint32 `gorm:"column:function_id" form:"function_id"`
	RuleId     uint32 `gorm:"column:rule_id" form:"rule_id"`
}

type UpdateAggregatorRequest struct {
	Id         uint32 `gorm:"primaryKey;column:id" form:"id"`
	Name       string `gorm:"column:name" form:"name"`
	Metric     string `gorm:"column:metric" form:"metric"`
	FunctionId uint32 `gorm:"column:function_id" form:"function_id"`
	RuleId     uint32 `gorm:"column:rule_id" form:"rule_id"`
}

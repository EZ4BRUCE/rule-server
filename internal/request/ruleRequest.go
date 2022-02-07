package request

// 接收聚合器的POST请求参数定义的结构体
// 其他的请求只有一个path参数（ID）或者没有参数（list）故不用构建请求结构体

type CreateRuleRequest struct {
	Level       string `gorm:"level" form:"level"`
	Action      string `gorm:"action" form:"action"`
	Description string `gorm:"description" form:"description"`
}

type UpdateRuleRequest struct {
	Id          uint32 `gorm:"primaryKey;column:id" form:"id"`
	Level       string `gorm:"level" form:"level"`
	Action      string `gorm:"action" form:"action"`
	Description string `gorm:"description" form:"description"`
}

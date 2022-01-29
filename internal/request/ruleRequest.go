package request

type CreateRuleRequest struct {
	Level       string `gorm:"level" form:"level"`
	Action      string `gorm:"action" form:"action"`
	Description string `gorm:"description" form:"description"`
}

// 其他的请求只有一个参数（ID）或者没有参数（list）故不用构建请求体

type UpdateRuleRequest struct {
	Id          uint32 `gorm:"primaryKey;column:id" form:"id"`
	Level       string `gorm:"level" form:"level"`
	Action      string `gorm:"action" form:"action"`
	Description string `gorm:"description" form:"description"`
}

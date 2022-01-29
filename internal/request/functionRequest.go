package request

type CreateFunctionRequest struct {
	Type        string  `gorm:"column:type" form:"type"`
	Threshold   float64 `gorm:"column:threshold" form:"threshold"`
	Description string  `gorm:"description" form:"description"`
}

// 其他的请求只有一个参数（ID）或者没有参数（list）故不用构建请求体

type UpdateFunctionRequest struct {
	Id          uint32  `gorm:"primaryKey;column:id" form:"id"`
	Type        string  `gorm:"column:type" form:"type"`
	Threshold   float64 `gorm:"column:threshold" form:"threshold"`
	Description string  `gorm:"description" form:"description"`
}

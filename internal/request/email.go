package request

// 接收聚合器的POST请求参数定义的结构体
// 其他的请求只有一个path参数（ID）或者没有参数（list）故不用构建请求结构体

type CreateEmailRequest struct {
	Address string `gorm:"address" form:"address"`
}

type UpdateEmailRequest struct {
	Id      uint32 `gorm:"primaryKey;column:id" form:"id"`
	Address string `gorm:"address" form:"address"`
}

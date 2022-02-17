package request

// 接收聚合器的POST请求参数定义的结构体
// 其他的请求只有一个path参数（ID）或者没有参数（list）故不用构建请求结构体

type CreateSmtpRequest struct {
	Host     string `gorm:"column:host" form:"host"`
	Port     int    `gorm:"column:port" form:"port"`
	IsSSL    bool   `gorm:"column:isSSL" form:"isSSL"`
	UserName string `gorm:"column:userName" form:"userName"`
	Password string `gorm:"column:password" form:"password"`
	From     string `gorm:"column:from" form:"from"`
}

type UpdateSmtpRequest struct {
	Id       uint32 `gorm:"primaryKey;column:id" form:"id"`
	Host     string `gorm:"column:host" form:"host"`
	Port     int    `gorm:"column:port" form:"port"`
	IsSSL    bool   `gorm:"column:isSSL" form:"isSSL"`
	UserName string `gorm:"column:userName" form:"userName"`
	Password string `gorm:"column:password" form:"password"`
	From     string `gorm:"column:from" form:"from"`
}

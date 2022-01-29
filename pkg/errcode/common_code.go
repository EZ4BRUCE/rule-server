package errcode

var (
	Success       = NewError(0, "操作成功")
	ServerError   = NewError(10000000, "服务内部错误")
	InvalidParams = NewError(10000001, "入参错误")
	NotFound      = NewError(10000002, "找不到满足条件的信息")
)

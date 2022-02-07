package app

import (
	"net/http"

	"github.com/EZ4BRUCE/rule-server/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// 预先封装返回消息体的函数

// response封装一个gin上下文
type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// 成功返回
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

// 错误返回
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}

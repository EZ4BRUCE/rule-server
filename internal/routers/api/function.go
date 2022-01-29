package api

import (
	"github.com/EZ4BRUCE/rule-server/internal/request"
	"github.com/EZ4BRUCE/rule-server/internal/service"
	"github.com/EZ4BRUCE/rule-server/pkg/app"
	"github.com/EZ4BRUCE/rule-server/pkg/convert"
	"github.com/EZ4BRUCE/rule-server/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Function struct{}

// 空结构体，只是给接口方法分类，或者用函数也可以
func NewFunction() Function {
	return Function{}
}

func (f Function) Create(c *gin.Context) {
	param := request.CreateFunctionRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.CreateFunction(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "创建成功"})

}

func (f Function) Delete(c *gin.Context) {
	response := app.NewResponse(c)
	uintId, err := convert.StrTo(c.Param("id")).UInt32()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.DeleteFunction(uintId)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("删除错误：" + err.Error()))
		return
	}

	response.ToResponse(gin.H{"msg": "删除成功"})

}

func (f Function) Update(c *gin.Context) {
	param := request.UpdateFunctionRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.UpdateFunction(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("更新错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "更新成功"})
}

func (f Function) Get(c *gin.Context) {
	response := app.NewResponse(c)
	uintId, err := convert.StrTo(c.Param("id")).UInt32()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	function, err := svc.GetFunction(uintId)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"Function": function, "msg": "获取成功"})
}

func (f Function) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	functions, err := svc.ListFunction()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"All functions:": functions, "msg": "获取成功"})

}

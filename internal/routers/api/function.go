package api

import (
	"github.com/EZ4BRUCE/rule-server/internal/request"
	"github.com/EZ4BRUCE/rule-server/internal/service"
	"github.com/EZ4BRUCE/rule-server/pkg/app"
	"github.com/EZ4BRUCE/rule-server/pkg/convert"
	"github.com/EZ4BRUCE/rule-server/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// 空结构体给接口方法分类
type Function struct{}

func NewFunction() Function {
	return Function{}
}

// swagger注释

// @Summary 新增聚合函数
// @Produce  json
// @Param type body string false "函数类型"
// @Param threshold body float32 false "阈值"
// @Param description body string false "描述"
// @Param 测试 body request.CreateFunctionRequest true "新建聚合函数测试"
// @Success 200 {object} model.Function "创建成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/function [post]
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

// @Summary 删除聚合函数
// @Produce  json
// @Param id path uint32 true "聚合函数 ID"
// @Success 200 {string} string "删除成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/function/{id} [delete]
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

// @Summary 更新聚合函数
// @Produce  json
// @Param id body uint32 false "聚合函数 ID"
// @Param type body string false "函数类型"
// @Param threshold body float32 false "阈值"
// @Param description body string false "描述"
// @Param 测试 body request.UpdateFunctionRequest true "更新聚合函数测试"
// @Success 200 {string} string "更新成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/function [put]
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

// @Summary 获取单个聚合函数
// @Produce  json
// @Param id path uint32 true "聚合函数 ID"
// @Success 200 {object} model.Function "获取成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/function/{id} [get]
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

// @Summary 获取所有聚合函数
// @Produce  json
// @Success 200 {array} model.Function "获取成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/functions [get]
func (f Function) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	functions, err := svc.ListFunctions()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"All functions:": functions, "msg": "获取成功"})

}

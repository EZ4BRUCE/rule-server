package api

import (
	"fmt"

	"github.com/EZ4BRUCE/rule-server/internal/request"
	"github.com/EZ4BRUCE/rule-server/internal/service"
	"github.com/EZ4BRUCE/rule-server/pkg/app"
	"github.com/EZ4BRUCE/rule-server/pkg/convert"
	"github.com/EZ4BRUCE/rule-server/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Aggregator struct{}

// 空结构体，只是给接口方法分类，或者用函数也可以
func NewAggregator() Aggregator {
	return Aggregator{}
}

// @Summary 新增聚合器
// @Produce  json
// @Param name body string true "名称"
// @Param metric body string true "监控指标"
// @Param function_id body int true "聚合函数"
// @Param rule_id body int true "告警规则"
// @Success 200 {object} model.Aggregator "创建成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/aggregator [post]
func (a Aggregator) Create(c *gin.Context) {
	param := request.CreateAggregatorRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	fmt.Println(param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.CreateAggregator(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "创建成功"})

}

// @Summary 删除聚合器
// @Produce  json
// @Param id path int true "聚合器 ID"
// @Success 200 {string} string "删除成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/aggregator/{id} [delete]
func (a Aggregator) Delete(c *gin.Context) {
	response := app.NewResponse(c)
	uintId, err := convert.StrTo(c.Param("id")).UInt32()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.DeleteAggregator(uintId)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("删除错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "删除成功"})
}

// @Summary 更新聚合器
// @Produce  json
// @Param id body int true "聚合器 ID"
// @Param name body string true "名称"
// @Param metric body string true "监控指标"
// @Param function_id body int true "聚合函数"
// @Param rule_id body int true "告警规则"
// @Success 200 {string} string "更新成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/aggregator [put]
func (a Aggregator) Update(c *gin.Context) {
	param := request.UpdateAggregatorRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.UpdateAggregator(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("更新错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "更新成功"})
}

// @Summary 获取单个聚合器
// @Produce  json
// @Param id path int true "聚合器 ID"
// @Success 200 {object} model.Aggregator "获取成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/aggregator/{id} [get]
func (a Aggregator) Get(c *gin.Context) {
	response := app.NewResponse(c)
	uintId, err := convert.StrTo(c.Param("id")).UInt32()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	aggregator, err := svc.GetAggregator(uintId)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"Aggregator": aggregator, "msg": "获取成功"})
}

// @Summary 获取所有聚合器
// @Produce  json
// @Success 200 {array} model.Aggregator "获取成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/aggregators [get]
func (a Aggregator) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	aggregators, err := svc.ListAggregators()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"All aggregators:": aggregators, "msg": "获取成功"})

}

// @Summary 获取特定指标的聚合器
// @Produce  json
// @Param metric path string true "监控指标"
// @Success 200 {array} model.Aggregator "获取成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/aggregator/search/{metric} [get]
func (a Aggregator) Search(c *gin.Context) {
	response := app.NewResponse(c)
	metric := c.Param("metric")
	svc := service.New(c.Request.Context())
	aggregators, err := svc.SearchAggregators(metric)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"All aggregators:": aggregators, "msg": "获取成功"})
}

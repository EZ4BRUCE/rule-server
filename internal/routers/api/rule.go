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
type Rule struct{}

func NewRule() Rule {
	return Rule{}
}

// swagger注释

// @Summary 新增告警规则
// @Produce  json
// @Param level body string false "警报等级"
// @Param action body string false "告警行为"
// @Param description body string false "描述"
// @Param 测试 body request.CreateRuleRequest true "新建告警规则测试"
// @Success 200 {object} model.Rule "创建成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/rule [post]
func (r Rule) Create(c *gin.Context) {
	param := request.CreateRuleRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.CreateRule(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "创建成功"})

}

// @Summary 删除告警规则
// @Produce  json
// @Param id path uint32 true "告警规则 ID"
// @Success 200 {string} string "删除成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/rule/{id} [delete]
func (r Rule) Delete(c *gin.Context) {
	response := app.NewResponse(c)
	uintId, err := convert.StrTo(c.Param("id")).UInt32()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.DeleteRule(uintId)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("删除错误：" + err.Error()))
		return
	}

	response.ToResponse(gin.H{"msg": "删除成功"})

}

// @Summary 更新告警规则
// @Produce  json
// @Param id body uint32 false "告警规则 ID"
// @Param level body string false "警报等级"
// @Param action body string false "告警行为"
// @Param description body string false "描述"
// @Param 测试 body request.UpdateRuleRequest true "更新告警规则测试"
// @Success 200 {string} string "更新成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/rule [put]
func (r Rule) Update(c *gin.Context) {
	param := request.UpdateRuleRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.UpdateRule(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("更新错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "更新成功"})
}

// @Summary 获取单个告警规则
// @Produce  json
// @Param id path uint32 true "告警规则 ID"
// @Success 200 {object} model.Rule "获取成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/rule/{id} [get]
func (r Rule) Get(c *gin.Context) {
	response := app.NewResponse(c)
	uintId, err := convert.StrTo(c.Param("id")).UInt32()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	rule, err := svc.GetRule(uintId)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"Rule": rule, "msg": "获取成功"})
}

// @Summary 获取所有告警规则
// @Produce  json
// @Success 200 {array} model.Rule "获取成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/rules [get]
func (r Rule) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	rules, err := svc.ListRules()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"All rules:": rules, "msg": "获取成功"})
}

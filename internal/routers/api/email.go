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
type Email struct{}

func NewEmail() Email {
	return Email{}
}

// swagger注释中为了方便说明，用了多个body，只有最后一个测试body会起作用

// @Summary 新增告警邮箱
// @Produce  json
// @Param address body string false "告警邮箱地址"
// @Param 测试 body request.CreateEmailRequest true "新建告警邮箱测试"
// @Success 200 {object} model.Email "创建成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/email [post]
func (r Email) Create(c *gin.Context) {
	param := request.CreateEmailRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.CreateEmail(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "创建成功"})

}

// @Summary 删除告警邮箱
// @Produce  json
// @Param id path uint32 true "告警邮箱 ID"
// @Success 200 {string} string "删除成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/email/{id} [delete]
func (r Email) Delete(c *gin.Context) {
	response := app.NewResponse(c)
	uintId, err := convert.StrTo(c.Param("id")).UInt32()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.DeleteEmail(uintId)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("删除错误：" + err.Error()))
		return
	}

	response.ToResponse(gin.H{"msg": "删除成功"})

}

// @Summary 更新告警邮箱
// @Produce  json
// @Param id body uint32 false "告警邮箱 ID"
// @Param address body string false "告警邮箱地址"
// @Param 测试 body request.UpdateEmailRequest true "更新告警邮箱测试"
// @Success 200 {string} string "更新成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/email [put]
func (r Email) Update(c *gin.Context) {
	param := request.UpdateEmailRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.UpdateEmail(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("更新错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "更新成功"})
}

// @Summary 获取单个告警邮箱
// @Produce  json
// @Param id path uint32 true "告警邮箱 ID"
// @Success 200 {object} model.Email "获取成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/email/{id} [get]
func (r Email) Get(c *gin.Context) {
	response := app.NewResponse(c)
	uintId, err := convert.StrTo(c.Param("id")).UInt32()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	email, err := svc.GetEmail(uintId)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"Email": email, "msg": "获取成功"})
}

// @Summary 获取所有告警邮箱
// @Produce  json
// @Success 200 {array} model.Email "获取成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/emails [get]
func (r Email) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	emails, err := svc.ListEmails()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"All emails:": emails, "msg": "获取成功"})
}

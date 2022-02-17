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
type Smtp struct{}

func NewSmtp() Smtp {
	return Smtp{}
}

// swagger注释中为了方便说明，用了多个body，只有最后一个测试body会起作用

// @Summary 新增SMTP邮箱
// @Produce  json
// @Param host body string false "SMTP服务器"
// @Param port body int false "端口"
// @Param isSSL body bool false "是否使用SSL加密"
// @Param userName body string false "SMTP登录名"
// @Param password body string false "SMTP密码"
// @Param form body string false "发送邮箱地址"
// @Param 测试 body request.CreateSmtpRequest true "新建SMTP邮箱测试"
// @Success 200 {object} model.Smtp "创建成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/smtp [post]
func (r Smtp) Create(c *gin.Context) {
	param := request.CreateSmtpRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.CreateSmtp(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "创建成功"})

}

// @Summary 删除SMTP邮箱
// @Produce  json
// @Param id path uint32 true "SMTP邮箱 ID"
// @Success 200 {string} string "删除成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/smtp/{id} [delete]
func (r Smtp) Delete(c *gin.Context) {
	response := app.NewResponse(c)
	uintId, err := convert.StrTo(c.Param("id")).UInt32()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.DeleteSmtp(uintId)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("删除错误：" + err.Error()))
		return
	}

	response.ToResponse(gin.H{"msg": "删除成功"})

}

// @Summary 更新SMTP邮箱
// @Produce  json
// @Param id body uint32 false "SMTP邮箱 ID"
// @Param host body string false "SMTP服务器"
// @Param port body int false "端口"
// @Param isSSL body bool false "是否使用SSL加密"
// @Param userName body string false "SMTP登录名"
// @Param password body string false "SMTP密码"
// @Param form body string false "发送邮箱地址"// @Param 测试 body request.UpdateSmtpRequest true "更新SMTP邮箱测试"
// @Success 200 {string} string "更新成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/smtp [put]
func (r Smtp) Update(c *gin.Context) {
	param := request.UpdateSmtpRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.UpdateSmtp(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("更新错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "更新成功"})
}

// @Summary 获取单个SMTP邮箱
// @Produce  json
// @Param id path uint32 true "SMTP邮箱 ID"
// @Success 200 {object} model.Smtp "获取成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/smtp/{id} [get]
func (r Smtp) Get(c *gin.Context) {
	response := app.NewResponse(c)
	uintId, err := convert.StrTo(c.Param("id")).UInt32()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("系统错误：" + err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	smtp, err := svc.GetSmtp(uintId)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"Smtp": smtp, "msg": "获取成功"})
}

// @Summary 获取所有SMTP邮箱
// @Produce  json
// @Success 200 {array} model.Smtp "获取成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/smtps [get]
func (r Smtp) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	smtps, err := svc.ListSmtps()
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails("获取错误：" + err.Error()))
		return
	}
	response.ToResponse(gin.H{"All smtps:": smtps, "msg": "获取成功"})
}

package routers

import (
	"github.com/EZ4BRUCE/rule-server/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	rule := api.NewRule()
	api := r.Group("/api/")
	{
		api.GET("rule/hello", rule.Hello)
	}
	return r
}

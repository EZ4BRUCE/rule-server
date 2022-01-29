package routers

import (
	"github.com/EZ4BRUCE/rule-server/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	aggregator := api.NewAggregator()
	function := api.NewFunction()
	rule := api.NewRule()
	api := r.Group("/api/")
	{
		api.POST("/aggregator", aggregator.Create)
		api.DELETE("/aggregator/:id", aggregator.Delete)
		api.PUT("/aggregator", aggregator.Update)
		api.GET("/aggregator/:id", aggregator.Get)
		api.GET("/aggregator/search/:metric", aggregator.Search)
		api.GET("/aggregators", aggregator.List)

		api.POST("/function", function.Create)
		api.DELETE("/function/:id", function.Delete)
		api.PUT("/function", function.Update)
		api.GET("/function/:id", function.Get)
		api.GET("/functions", function.List)

		api.POST("/rule", rule.Create)
		api.DELETE("/rule/:id", rule.Delete)
		api.PUT("/rule", rule.Update)
		api.GET("/rule/:id", rule.Get)
		api.GET("/rules", rule.List)

	}
	return r
}

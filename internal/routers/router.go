package routers

import (
	_ "github.com/EZ4BRUCE/rule-server/docs"
	"github.com/EZ4BRUCE/rule-server/internal/routers/api"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// localhost:1016/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	aggregator := api.NewAggregator()
	function := api.NewFunction()
	rule := api.NewRule()
	api := r.Group("/api/")
	{
		api.POST("/aggregator", aggregator.Create)
		api.DELETE("/aggregator/:id", aggregator.Delete)
		api.PUT("/aggregator", aggregator.Update)
		api.GET("/aggregator/:id", aggregator.Get)
		api.GET("/aggregators", aggregator.List)
		api.GET("/aggregator/search/:metric", aggregator.Search)

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

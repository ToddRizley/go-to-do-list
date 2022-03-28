package routes

import (
	"github.com/gin-gonic/gin"
	"go-to-do-list/controllers"
)

func BuildRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("todo", controllers.GetTodoItems)
		api.GET("todo/:id", controllers.GetATodoItem)
		api.GET("healthz", controllers.GetAPIHealth)
		api.POST("todo", controllers.CreateATodoItem)
		api.PUT("todo/:id", controllers.UpdateATodoItem)
		api.DELETE("todo/:id", controllers.DeleteATodoItem)
	}
	return r
}

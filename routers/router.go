package routers

import (
	"github.com/gin-gonic/gin"

	"gin_video_snapshot/controller"
)

func InitRouter() *gin.Engine {
	g := gin.Default()
	g.Static("/static", "static")
	g.LoadHTMLGlob("templates/*")

	g.GET("/", controller.Index)

	v1Group := g.Group("/v1")
	{
		// 添加代办事项
		v1Group.POST("/todo", controller.CreateTodo)

		// 获取全部代办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 更新
		v1Group.PUT("/todo/:id", controller.UpdateTodo)

		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return g
}

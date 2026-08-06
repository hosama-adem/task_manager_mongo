package router

import (
	"Task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:id", controllers.GetTaskByID)
	router.PUT("/tasks/:id", controllers.UpdateTaskByID)
	router.DELETE("/tasks/:id", controllers.RemoveTaskByID)
	router.POST("/tasks", controllers.AddTask)

	return router
}

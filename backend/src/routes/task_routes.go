package routes

import (
	task_handler "backend/src/handlers/task"

	"github.com/labstack/echo/v4"
)

func SetupTaskRoutes(e *echo.Echo, taskHandler *task_handler.TaskHandler) {
	taskGroup := e.Group("/tasks")
	taskGroup.GET("", taskHandler.HandlerGetAllTasks)
	taskGroup.GET("/:id", taskHandler.HandlerGetTaskByID)
	taskGroup.POST("", taskHandler.HandlerCreateTask)
	taskGroup.PUT("/:id", taskHandler.HandlerUpdateTask)
	taskGroup.DELETE("/:id", taskHandler.HandlerDeleteTask)
}

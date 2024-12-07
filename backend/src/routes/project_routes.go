package routes

import (
	project_handler "backend/src/handlers/project"

	"github.com/labstack/echo/v4"
)

func SetupProjectRoutes(e *echo.Echo, projectHandler *project_handler.ProjectHandler) {
	projectGroup := e.Group("/projects")
	projectGroup.GET("", projectHandler.HandlerGetAllProjects)
	projectGroup.GET("/:id", projectHandler.HandlerGetProjectByID)
	projectGroup.POST("", projectHandler.HandlerCreateProject)
	projectGroup.PUT("/:id", projectHandler.HandlerUpdateProject)
	projectGroup.DELETE("/:id", projectHandler.HandlerDeleteProject)
}

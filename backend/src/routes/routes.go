package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	sample_handler "backend/src/handlers/sample"
	sample_repository "backend/src/repositories/sample"
	sample_service "backend/src/services/sample"

	auth_handler "backend/src/handlers/auth"
	user_repository "backend/src/repositories/user"
	auth_service "backend/src/services/auth"

	project_handler "backend/src/handlers/project"
	project_repository "backend/src/repositories/project"
	project_service "backend/src/services/project"

	task_handler "backend/src/handlers/task"
	task_repository "backend/src/repositories/task"
	task_service "backend/src/services/task"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ルートのセットアップ
func RoutesSetUp(e *echo.Echo, db *sql.DB, dbGorm *gorm.DB) {
	fmt.Println("Setting up routes...")

	// データベース接続確認
	e.GET("/", func(c echo.Context) error {
		fmt.Println("GET / connection check.")
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM samples").Scan(&count)
		if err != nil {
			fmt.Println("Error querying database:", err)
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error querying database: %v", err))
		}
		fmt.Println("Connection OK.Row count:", count)
		return c.String(http.StatusOK, fmt.Sprintf("Row count: %d", count))
	})

	// DI
	sampleRepo := sample_repository.NewSampleRepository(dbGorm)
	sampleService := sample_service.NewSampleService(sampleRepo)
	sampleHandler := sample_handler.NewSampleHandler(sampleService)

	authRepo := user_repository.NewUserRepository(dbGorm)
	authService := auth_service.NewAuthUserService(authRepo)
	authHandler := auth_handler.NewAuthUserHandler(authService)

	projectRepo := project_repository.NewProjectRepository(dbGorm)
	projectService := project_service.NewProjectService(projectRepo)
	projectHandler := project_handler.NewProjectHandler(projectService)

	taskRepo := task_repository.NewTaskRepository(dbGorm)
	taskService := task_service.NewTaskService(taskRepo)
	taskHandler := task_handler.NewTaskHandler(taskService)

	// サンプルルートのセットアップ
	SetupSampleRoutes(e, sampleHandler)

	// 認証ルートのセットアップ
	SetupAuthUserRoutes(e, authHandler)

	// プロジェクトルートのセットアップ
	SetupProjectRoutes(e, projectHandler)

	// タスクルートのセットアップ
	SetupTaskRoutes(e, taskHandler)

	fmt.Println("Routes set up")
}

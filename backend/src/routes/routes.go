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

	// サンプルルートのセットアップ
	SetupSampleRoutes(e, sampleHandler)

	// 認証ルートのセットアップ
	SetupAuthUserRoutes(e, authHandler)

	fmt.Println("Routes set up")
}

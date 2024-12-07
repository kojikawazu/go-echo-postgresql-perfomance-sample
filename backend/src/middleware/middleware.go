package middleware

import "github.com/labstack/echo/v4"

// SetUpMiddleware はミドルウェアを設定します
func SetUpMiddleware(e *echo.Echo) {
	// 実行時間計測ミドルウェア
	e.Use(MeasureExecutionTime)
}

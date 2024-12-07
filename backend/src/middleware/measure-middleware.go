package middleware

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// MeasureExecutionTime ミドルウェアでAPI全体の実行時間を計測
func MeasureExecutionTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエスト開始時刻を記録
		start := time.Now()

		// 次のハンドラーを実行
		err := next(c)

		// 実行時間を計測
		duration := time.Since(start)

		// 実行時間をログに記録
		log.Printf("[%s] %s %s took %v", c.Request().Method, c.Request().Host, c.Request().RequestURI, duration)

		return err
	}
}

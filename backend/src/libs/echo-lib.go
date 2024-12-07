package lib

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// Echoインスタンスの作成
func EchoSetUp() (*echo.Echo, error) {
	fmt.Println("Creating Echo instance...")

	e := echo.New()

	fmt.Println("Echo instance created")
	return e, nil
}

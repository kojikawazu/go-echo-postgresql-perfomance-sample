package routes

import (
	auth_handler "backend/src/handlers/auth"

	"github.com/labstack/echo/v4"
)

func SetupAuthUserRoutes(e *echo.Echo, authHandler *auth_handler.AuthUserHandler) {
	authGroup := e.Group("/auth")
	authGroup.POST("/signin", authHandler.SignIn)
	authGroup.POST("/signup", authHandler.SignUp)
	authGroup.POST("/signout", authHandler.SignOut)
	authGroup.GET("/user", authHandler.GetAuthenticatedUser)
}

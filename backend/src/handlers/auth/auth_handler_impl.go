package auth_handler

import auth_service "backend/src/services/auth"

type AuthUserHandler struct {
	service auth_service.AuthUserService
}

// NewAuthUserHandler はDI対応のハンドラーを返します
func NewAuthUserHandler(service auth_service.AuthUserService) *AuthUserHandler {
	return &AuthUserHandler{service: service}
}

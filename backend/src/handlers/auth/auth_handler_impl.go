package auth_handler

import auth_service "backend/src/services/auth"

// 認証ユーザーハンドラ
type AuthUserHandler struct {
	service auth_service.AuthUserService
}

// コンストラクタ
func NewAuthUserHandler(service auth_service.AuthUserService) *AuthUserHandler {
	return &AuthUserHandler{service}
}

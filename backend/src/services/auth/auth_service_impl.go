package auth_user_service

import (
	auth_user_model "backend/src/models/user"
	user_repository "backend/src/repositories/user"
)

// AuthUserService インターフェース
type AuthUserService interface {
	// ユーザー認証
	ServiceAuthenticateUser(email, password string) (string, error)
	// ユーザー登録
	ServiceRegisterUser(username, email, password string) (string, error)
	// ユーザー情報取得
	ServiceGetUserByID(userID string) (*auth_user_model.User, error)
}

// authUserService 構造体
type authUserService struct {
	repo user_repository.UserRepository
}

// NewUserService はDI対応のサービスを返します
func NewAuthUserService(repo user_repository.UserRepository) AuthUserService {
	return &authUserService{repo: repo}
}

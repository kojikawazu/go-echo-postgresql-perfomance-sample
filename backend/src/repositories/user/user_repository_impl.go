package user_repository

import (
	auth_user_model "backend/src/models/user"

	"gorm.io/gorm"
)

// UserRepository はユーザーリポジトリのインターフェイス
type UserRepository interface {
	// メールアドレスをキーにユーザーを取得
	RepositoryGetUserByEmail(email string) (*auth_user_model.User, error)
	// ユーザーを取得
	RepositoryGetUserByID(userID string) (*auth_user_model.User, error)
	// ユーザーを作成
	RepositoryCreateUser(user *auth_user_model.User) error
}

// userRepository はユーザーリポジトリの実装
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository はDI対応のリポジトリを返します
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

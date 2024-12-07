package auth_user_model

import (
	"time"

	"gorm.io/gorm"
)

// User 構造体
type User struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username  string `gorm:"size:100;unique;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Password  string `gorm:"size:255;not null"` // ハッシュ化されたパスワードを保存
	CreatedAt time.Time
	UpdatedAt time.Time
}

// AutoMigrate マイグレーション用の関数
func AutoMigrate(db *gorm.DB) error {
	// UUID生成拡張を有効化（PostgreSQL専用）
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		return err
	}

	// テーブルのマイグレーション
	return db.AutoMigrate(&User{})
}

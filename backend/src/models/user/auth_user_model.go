package auth_user_model

import (
	"time"
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

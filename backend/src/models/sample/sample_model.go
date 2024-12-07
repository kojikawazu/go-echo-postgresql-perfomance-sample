package sample_model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Sample モデル
type Sample struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"` // UUIDを使用
	Name      string `gorm:"size:100"`                                        // 最大100文字
	Value     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// マイグレーション用の関数
func AutoMigrate(db *gorm.DB) error {
	fmt.Println("AutoMigrate start...")

	// PostgreSQLの場合、uuid_generate_v4()関数を有効にする必要があります
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		fmt.Println("AutoMigrate error:", err)
		return err
	}

	fmt.Println("AutoMigrate end...")
	return db.AutoMigrate(&Sample{})
}

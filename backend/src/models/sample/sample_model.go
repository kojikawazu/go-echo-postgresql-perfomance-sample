package sample_model

import (
	"time"
)

// Sample モデル
type Sample struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"` // UUIDを使用
	Name      string `gorm:"size:100"`                                        // 最大100文字
	Value     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

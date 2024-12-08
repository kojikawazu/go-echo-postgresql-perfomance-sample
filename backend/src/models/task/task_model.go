package task_model

import (
	"time"

	user_model "backend/src/models/user"
)

// Task 構造体
type Task struct {
	ID          string           `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProjectID   string           `gorm:"type:uuid;not null;index:idx_project_status"` // プロジェクトIDとのリレーション
	Name        string           `gorm:"size:255;not null;index"`
	Description string           `gorm:"type:text;"`
	Status      string           `gorm:"size:50;default:'pending';index:idx_project_status"`
	AssignedTo  *string          `gorm:"type:uuid;default:null;index"` // ユーザーIDへのリレーションを想定
	User        *user_model.User `gorm:"foreignKey:AssignedTo"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

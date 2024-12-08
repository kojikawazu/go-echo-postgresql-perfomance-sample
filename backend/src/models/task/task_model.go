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

// TaskWithUser 構造体
type TaskWithUser struct {
	ID            string     `gorm:"column:id"`
	ProjectID     string     `gorm:"column:project_id"`
	Name          string     `gorm:"column:name"`
	Description   string     `gorm:"column:description"`
	Status        string     `gorm:"column:status"`
	AssignedTo    *string    `gorm:"column:assigned_to"`
	CreatedAt     time.Time  `gorm:"column:created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at"`
	UserID        *string    `gorm:"column:user_id"`
	UserUserName  *string    `gorm:"column:user_username"`
	UserEmail     *string    `gorm:"column:user_email"`
	UserCreatedAt *time.Time `gorm:"column:user_created_at"`
	UserUpdatedAt *time.Time `gorm:"column:user_updated_at"`
}

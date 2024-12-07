package project_model

import (
	"time"

	task_model "backend/src/models/task"
)

// Project 構造体
type Project struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Tasks       []task_model.Task `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}

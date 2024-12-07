package task_repository

import (
	task_model "backend/src/models/task"

	"gorm.io/gorm"
)

// TaskRepository インターフェース
type TaskRepository interface {
	RepositoryGetAll() ([]task_model.Task, error)
	RepositoryGetByID(id string) (*task_model.Task, error)
	RepositoryCreate(task *task_model.Task) error
	RepositoryUpdate(task *task_model.Task) error
	RepositoryDelete(id string) error
}

// taskRepository 構造体
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository 関数
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

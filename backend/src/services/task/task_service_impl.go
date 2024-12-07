package task_service

import (
	task_model "backend/src/models/task"
	task_repository "backend/src/repositories/task"
)

// TaskService インターフェース
type TaskService interface {
	ServiceGetAllTasks() ([]task_model.Task, error)
	ServiceGetTaskByID(id string) (*task_model.Task, error)
	ServiceCreateTask(task *task_model.Task) error
	ServiceUpdateTask(task *task_model.Task) error
	ServiceDeleteTask(id string) error
}

// taskService 構造体
type taskService struct {
	repo task_repository.TaskRepository
}

// NewTaskService 関数
func NewTaskService(repo task_repository.TaskRepository) TaskService {
	return &taskService{repo}
}

package task_handler

import task_service "backend/src/services/task"

// タスクハンドラ
type TaskHandler struct {
	service task_service.TaskService
}

// コンストラクタ
func NewTaskHandler(service task_service.TaskService) *TaskHandler {
	return &TaskHandler{service}
}

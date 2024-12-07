package task_service

import (
	task_model "backend/src/models/task"
	logging_utils "backend/src/utils/logging"
	"errors"
)

// 全権取得
func (s *taskService) ServiceGetAllTasks() ([]task_model.Task, error) {
	start := logging_utils.LogStart()

	tasks, err := s.repo.RepositoryGetAll()
	if err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return tasks, nil
}

// ID取得
func (s *taskService) ServiceGetTaskByID(id string) (*task_model.Task, error) {
	start := logging_utils.LogStart()

	// バリデーション
	if id == "" {
		logging_utils.LogError(start, errors.New("id is required"))
		return nil, errors.New("id is required")
	}

	task, err := s.repo.RepositoryGetByID(id)
	if err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return task, nil
}

// 作成
func (s *taskService) ServiceCreateTask(task *task_model.Task) error {
	start := logging_utils.LogStart()

	// バリデーション
	if task.ProjectID == "" {
		logging_utils.LogError(start, errors.New("project_id is required"))
		return errors.New("project_id is required")
	}
	if task.Name == "" {
		logging_utils.LogError(start, errors.New("name is required"))
		return errors.New("name is required")
	}
	if task.Description == "" {
		logging_utils.LogError(start, errors.New("description is required"))
		return errors.New("description is required")
	}
	if task.Status == "" {
		logging_utils.LogError(start, errors.New("status is required"))
		return errors.New("status is required")
	}

	err := s.repo.RepositoryCreate(task)
	if err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

// 更新
func (s *taskService) ServiceUpdateTask(task *task_model.Task) error {
	start := logging_utils.LogStart()

	// バリデーション
	if task.ID == "" {
		logging_utils.LogError(start, errors.New("id is required"))
		return errors.New("id is required")
	}
	if task.ProjectID == "" {
		logging_utils.LogError(start, errors.New("project_id is required"))
		return errors.New("project_id is required")
	}
	if task.Name == "" {
		logging_utils.LogError(start, errors.New("name is required"))
		return errors.New("name is required")
	}
	if task.Description == "" {
		logging_utils.LogError(start, errors.New("description is required"))
		return errors.New("description is required")
	}
	if task.Status == "" {
		logging_utils.LogError(start, errors.New("status is required"))
		return errors.New("status is required")
	}

	err := s.repo.RepositoryUpdate(task)
	if err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

// 削除
func (s *taskService) ServiceDeleteTask(id string) error {
	start := logging_utils.LogStart()

	// バリデーション
	if id == "" {
		logging_utils.LogError(start, errors.New("id is required"))
		return errors.New("id is required")
	}

	err := s.repo.RepositoryDelete(id)
	if err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

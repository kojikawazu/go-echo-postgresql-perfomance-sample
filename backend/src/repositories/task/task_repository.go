package task_repository

import (
	task_model "backend/src/models/task"
	logging_utils "backend/src/utils/logging"
)

// 全件取得
func (r *taskRepository) RepositoryGetAll() ([]task_model.Task, error) {
	start := logging_utils.LogStart()

	var tasks []task_model.Task
	// User を結合
	if err := r.db.
		Preload("User").
		Find(&tasks).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return tasks, nil
}

// ID取得
func (r *taskRepository) RepositoryGetByID(id string) (*task_model.Task, error) {
	start := logging_utils.LogStart()

	var task task_model.Task
	// User を結合
	if err := r.db.
		Preload("User").
		First(&task, "id = ?", id).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return &task, nil
}

// 作成
func (r *taskRepository) RepositoryCreate(task *task_model.Task) error {
	start := logging_utils.LogStart()

	// トランザクション開始
	tx := r.db.Begin()
	defer func() {
		// ロールバック
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 作成
	if err := tx.Create(task).Error; err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	// コミット
	if err := tx.Commit().Error; err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

// 更新
func (r *taskRepository) RepositoryUpdate(task *task_model.Task) error {
	start := logging_utils.LogStart()

	// トランザクション開始
	tx := r.db.Begin()
	defer func() {
		// ロールバック
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新
	if err := tx.Save(task).Error; err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	// コミット
	if err := tx.Commit().Error; err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

// 削除
func (r *taskRepository) RepositoryDelete(id string) error {
	start := logging_utils.LogStart()

	// トランザクション開始
	tx := r.db.Begin()
	defer func() {
		// ロールバック
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 削除
	if err := tx.Delete(&task_model.Task{}, "id = ?", id).Error; err != nil {
		logging_utils.LogError(start, err)
		tx.Rollback()
		return err
	}

	// コミット
	if err := tx.Commit().Error; err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

package task_repository

import (
	task_model "backend/src/models/task"
	user_model "backend/src/models/user"
	logging_utils "backend/src/utils/logging"
	"errors"
	"fmt"
	"time"
)

// 全件取得
func (r *taskRepository) RepositoryGetAll() ([]task_model.Task, error) {
	start := logging_utils.LogStart()

	var rows []task_model.TaskWithUser
	if err := r.db.Raw(`
        SELECT
            tasks.id AS id,
			tasks.project_id AS project_id,
            tasks.name AS name,
            tasks.description AS description,
            tasks.status AS status,
			tasks.assigned_to AS assigned_to,
            tasks.created_at AS created_at,
            tasks.updated_at AS updated_at,
            users.id AS user_id,
            users.username AS user_username,
            users.email AS user_email,
            users.created_at AS user_created_at,
            users.updated_at AS user_updated_at
        FROM tasks
        LEFT JOIN users ON tasks.assigned_to = users.id
    `).Scan(&rows).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	// Goでタスクごとにグループ化
	taskMap := make(map[string]*task_model.Task)
	for _, row := range rows {
		// タスクが未登録なら初期化
		if _, exists := taskMap[row.ID]; !exists {
			taskMap[row.ID] = &task_model.Task{
				ID:          row.ID,
				ProjectID:   row.ProjectID,
				Name:        row.Name,
				Description: row.Description,
				Status:      row.Status,
				AssignedTo:  row.AssignedTo,
				CreatedAt:   row.CreatedAt,
				UpdatedAt:   row.UpdatedAt,
			}
		}

		// ユーザー情報がある場合のみ設定
		if row.UserID != nil && row.AssignedTo != nil {
			user := &user_model.User{
				ID:        *row.UserID,
				Username:  *row.UserUserName,
				Email:     *row.UserEmail,
				Password:  "", // セキュリティのため設定しない
				CreatedAt: *row.UserCreatedAt,
				UpdatedAt: *row.UserUpdatedAt,
			}
			taskMap[row.ID].User = user
		}
	}

	// Mapをスライスに変換
	var tasks []task_model.Task
	for _, task := range taskMap {
		tasks = append(tasks, *task)
	}

	logging_utils.LogEnd(start)
	return tasks, nil
}

// ID取得
func (r *taskRepository) RepositoryGetByID(id string) (*task_model.Task, error) {
	start := logging_utils.LogStart()

	var rows []task_model.TaskWithUser
	if err := r.db.Raw(`
		SELECT
			tasks.id AS id,
			tasks.project_id AS project_id,
			tasks.name AS name,
			tasks.description AS description,
			tasks.status AS status,
			tasks.assigned_to AS assigned_to,
			tasks.created_at AS created_at,
			tasks.updated_at AS updated_at,
			users.id AS user_id,
			users.username AS user_username,
			users.email AS user_email,
			users.created_at AS user_created_at,
			users.updated_at AS user_updated_at
		FROM tasks
		LEFT JOIN users ON tasks.assigned_to = users.id
		WHERE tasks.id = ?
	`, id).Scan(&rows).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	// データがない場合のエラーハンドリング
	if len(rows) == 0 {
		logging_utils.LogError(start, fmt.Errorf("task not found"))
		return nil, fmt.Errorf("task not found")
	}

	// タスクを初期化
	task := &task_model.Task{
		ID:          rows[0].ID,
		ProjectID:   rows[0].ProjectID,
		Name:        rows[0].Name,
		Description: rows[0].Description,
		Status:      rows[0].Status,
		AssignedTo:  rows[0].AssignedTo,
		CreatedAt:   rows[0].CreatedAt,
		UpdatedAt:   rows[0].UpdatedAt,
		User:        nil,
	}

	// タスクを追加
	for _, row := range rows {
		if row.UserID != nil && row.AssignedTo != nil {
			task.User = &user_model.User{
				ID:        *row.UserID,
				Username:  *row.UserUserName,
				Email:     *row.UserEmail,
				Password:  "", // セキュリティのため設定しない
				CreatedAt: *row.UserCreatedAt,
				UpdatedAt: *row.UserUpdatedAt,
			}
		}
	}

	logging_utils.LogEnd(start)
	return task, nil
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
	if err := tx.Model(&task).Updates(map[string]interface{}{
		"name":        task.Name,
		"description": task.Description,
		"status":      task.Status,
		"assigned_to": task.AssignedTo,
		"updated_at":  time.Now(),
	}).Error; err != nil {
		logging_utils.LogError(start, fmt.Errorf("failed to update project: %w", err))
		tx.Rollback()
		return err
	}

	// コミット
	if err := tx.Commit().Error; err != nil {
		logging_utils.LogError(start, fmt.Errorf("failed to commit transaction: %w", err))
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
	result := tx.Delete(&task_model.Task{}, "id = ?", id)
	if result.Error != nil {
		logging_utils.LogError(start, fmt.Errorf("failed to delete task: %w", result.Error))
		tx.Rollback()
		return result.Error
	}

	// 対象データが存在しない場合
	if result.RowsAffected == 0 {
		logging_utils.LogError(start, fmt.Errorf("task not found"))
		tx.Rollback()
		return errors.New("task not found")
	}

	// コミット
	if err := tx.Commit().Error; err != nil {
		logging_utils.LogError(start, fmt.Errorf("failed to commit transaction: %w", err))
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

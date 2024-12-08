package project_repository

import (
	project_model "backend/src/models/project"
	task_model "backend/src/models/task"
	logging_utils "backend/src/utils/logging"
	"errors"
	"fmt"
	"time"
)

// 全件取得
func (r *projectRepository) RepositoryGetAll() ([]project_model.Project, error) {
	start := logging_utils.LogStart()

	// 生データ取得
	var rows []project_model.ProjectWithTasks
	if err := r.db.Raw(`
		SELECT
			projects.id AS id,
			projects.name AS name,
			projects.description AS description,
			projects.created_at AS created_at,
			projects.updated_at AS updated_at,
			tasks.id AS task_id,
			tasks.name AS task_name,
			tasks.description AS task_description,
			tasks.status AS task_status,
			tasks.assigned_to AS task_assigned_to,
			tasks.created_at AS task_created_at,
			tasks.updated_at AS task_updated_at
		FROM projects
		LEFT JOIN tasks ON tasks.project_id = projects.id
	`).Scan(&rows).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	// Goでプロジェクトごとにグループ化
	projectMap := make(map[string]*project_model.Project)
	for _, row := range rows {
		// プロジェクトが未登録なら初期化
		if _, exists := projectMap[row.ID]; !exists {
			projectMap[row.ID] = &project_model.Project{
				ID:          row.ID,
				Name:        row.Name,
				Description: row.Description,
				CreatedAt:   row.CreatedAt,
				UpdatedAt:   row.UpdatedAt,
				Tasks:       []task_model.Task{},
			}
		}

		// タスクがある場合、追加
		if row.TaskID != nil {
			projectMap[row.ID].Tasks = append(projectMap[row.ID].Tasks, task_model.Task{
				ID:          *row.TaskID,
				Name:        *row.TaskName,
				Description: *row.TaskDescription,
				Status:      *row.TaskStatus,
				AssignedTo:  row.TaskAssignedTo,
				CreatedAt:   *row.TaskCreatedAt,
				UpdatedAt:   *row.TaskUpdatedAt,
			})
		}
	}

	// Mapをスライスに変換
	var projects []project_model.Project
	for _, project := range projectMap {
		projects = append(projects, *project)
	}

	logging_utils.LogEnd(start)
	return projects, nil
}

// ID取得
func (r *projectRepository) RepositoryGetByID(id string) (*project_model.Project, error) {
	start := logging_utils.LogStart()

	var rows []project_model.ProjectWithTasks
	if err := r.db.Raw(`
		SELECT
			projects.id AS id,
			projects.name AS name,
			projects.description AS description,
			projects.created_at AS created_at,
			projects.updated_at AS updated_at,
			tasks.id AS task_id,
			tasks.name AS task_name,
			tasks.description AS task_description,
			tasks.status AS task_status,
			tasks.assigned_to AS task_assigned_to,
			tasks.created_at AS task_created_at,
			tasks.updated_at AS task_updated_at
		FROM projects
		LEFT JOIN tasks ON tasks.project_id = projects.id
		WHERE projects.id = ?
	`, id).Scan(&rows).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	// データがない場合のエラーハンドリング
	if len(rows) == 0 {
		logging_utils.LogError(start, fmt.Errorf("project not found"))
		return nil, fmt.Errorf("project not found")
	}

	// プロジェクトを初期化
	project := &project_model.Project{
		ID:          rows[0].ID,
		Name:        rows[0].Name,
		Description: rows[0].Description,
		CreatedAt:   rows[0].CreatedAt,
		UpdatedAt:   rows[0].UpdatedAt,
		Tasks:       []task_model.Task{},
	}

	// タスクを追加
	for _, row := range rows {
		if row.TaskID != nil {
			project.Tasks = append(project.Tasks, task_model.Task{
				ID:          *row.TaskID,
				Name:        *row.TaskName,
				Description: *row.TaskDescription,
				Status:      *row.TaskStatus,
				AssignedTo:  row.TaskAssignedTo,
				CreatedAt:   *row.TaskCreatedAt,
				UpdatedAt:   *row.TaskUpdatedAt,
			})
		}
	}

	logging_utils.LogEnd(start)
	return project, nil
}

// 作成
func (r *projectRepository) RepositoryCreate(project *project_model.Project) error {
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
	if err := tx.Create(project).Error; err != nil {
		logging_utils.LogError(start, fmt.Errorf("failed to create project: %w", err))
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

// 更新
func (r *projectRepository) RepositoryUpdate(project *project_model.Project) error {
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
	if err := tx.Model(&project).Updates(map[string]interface{}{
		"name":        project.Name,
		"description": project.Description,
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
func (r *projectRepository) RepositoryDelete(id string) error {
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
	result := tx.Delete(&project_model.Project{}, "id = ?", id)
	if result.Error != nil {
		logging_utils.LogError(start, fmt.Errorf("failed to delete project: %w", result.Error))
		tx.Rollback()
		return result.Error
	}

	// 対象データが存在しない場合
	if result.RowsAffected == 0 {
		logging_utils.LogError(start, fmt.Errorf("project not found"))
		tx.Rollback()
		return errors.New("project not found")
	}

	// コミット
	if err := tx.Commit().Error; err != nil {
		logging_utils.LogError(start, fmt.Errorf("failed to commit transaction: %w", err))
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

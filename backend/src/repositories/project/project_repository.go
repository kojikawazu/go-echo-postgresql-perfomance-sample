package project_repository

import (
	project_model "backend/src/models/project"
	logging_utils "backend/src/utils/logging"
)

// 全件取得
func (r *projectRepository) RepositoryGetAll() ([]project_model.Project, error) {
	start := logging_utils.LogStart()

	var projects []project_model.Project
	// Tasks と Tasks 内の User を結合
	if err := r.db.
		Preload("Tasks").
		Preload("Tasks.User").
		Find(&projects).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return projects, nil
}

// ID取得
func (r *projectRepository) RepositoryGetByID(id string) (*project_model.Project, error) {
	start := logging_utils.LogStart()

	var project project_model.Project
	// Tasks と Tasks 内の User を結合
	if err := r.db.
		Preload("Tasks").
		Preload("Tasks.User").
		First(&project, "id = ?", id).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return &project, nil
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
	if err := tx.Save(project).Error; err != nil {
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
	if err := tx.Delete(&project_model.Project{}, "id = ?", id).Error; err != nil {
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

package sample_repository

import (
	sample_model "backend/src/models/sample"
	logging_utils "backend/src/utils/logging"
)

// 全取得
func (r *sampleRepository) RepositoryGetAll() ([]sample_model.Sample, error) {
	start := logging_utils.LogStart()

	var samples []sample_model.Sample
	if err := r.db.Find(&samples).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return samples, nil
}

// ID取得
func (r *sampleRepository) RepositoryGetByID(id string) (*sample_model.Sample, error) {
	start := logging_utils.LogStart()

	var sample sample_model.Sample
	if err := r.db.First(&sample, "id = ?", id).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return &sample, nil
}

// 作成
func (r *sampleRepository) RepositoryCreate(sample *sample_model.Sample) error {
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
	if err := tx.Create(sample).Error; err != nil {
		logging_utils.LogError(start, err)
		// ロールバック
		tx.Rollback()
		return err
	}

	// トランザクション終了
	if err := tx.Commit().Error; err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

// 更新
func (r *sampleRepository) RepositoryUpdate(sample *sample_model.Sample) error {
	start := logging_utils.LogStart()

	// 存在確認
	if _, err := r.RepositoryGetByID(sample.ID); err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	// トランザクション開始
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新
	if err := tx.Save(sample).Error; err != nil {
		logging_utils.LogError(start, err)
		// ロールバック
		tx.Rollback()
		return err
	}

	// トランザクション終了
	if err := tx.Commit().Error; err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

// 削除
func (r *sampleRepository) RepositoryDelete(id string) error {
	start := logging_utils.LogStart()

	// トランザクション開始
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 削除
	if err := tx.Delete(&sample_model.Sample{}, "id = ?", id).Error; err != nil {
		logging_utils.LogError(start, err)
		tx.Rollback()
		return err
	}

	// トランザクション終了
	if err := tx.Commit().Error; err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

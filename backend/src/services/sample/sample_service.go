package sample_service

import (
	sample_model "backend/src/models/sample"
	logging_utils "backend/src/utils/logging"
	"errors"
)

// 全取得
func (s *sampleService) ServiceGetAllSamples() ([]sample_model.Sample, error) {
	start := logging_utils.LogStart()

	result, err := s.repo.RepositoryGetAll()
	if err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return result, nil
}

// ID取得
func (s *sampleService) ServiceGetSampleByID(id string) (*sample_model.Sample, error) {
	start := logging_utils.LogStart()

	// バリデーション
	if id == "" {
		logging_utils.LogError(start, errors.New("id is empty"))
		return nil, errors.New("id is empty")
	}

	result, err := s.repo.RepositoryGetByID(id)
	if err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return result, nil
}

// 作成
func (s *sampleService) ServiceCreateSample(sample *sample_model.Sample) error {
	start := logging_utils.LogStart()

	// バリデーション
	if sample.Name == "" {
		logging_utils.LogError(start, errors.New("name is empty"))
		return errors.New("name is empty")
	}
	if sample.Value == 0 {
		logging_utils.LogError(start, errors.New("value is 0"))
		return errors.New("value is 0")
	}

	err := s.repo.RepositoryCreate(sample)
	if err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

// 更新
func (s *sampleService) ServiceUpdateSample(sample *sample_model.Sample) error {
	start := logging_utils.LogStart()

	// バリデーション
	if sample.ID == "" {
		logging_utils.LogError(start, errors.New("id is empty"))
		return errors.New("id is empty")
	}
	if sample.Name == "" {
		logging_utils.LogError(start, errors.New("name is empty"))
		return errors.New("name is empty")
	}
	if sample.Value == 0 {
		logging_utils.LogError(start, errors.New("value is 0"))
		return errors.New("value is 0")
	}

	err := s.repo.RepositoryUpdate(sample)
	if err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

// 削除
func (s *sampleService) ServiceDeleteSample(id string) error {
	start := logging_utils.LogStart()

	// バリデーション
	if id == "" {
		logging_utils.LogError(start, errors.New("id is empty"))
		return errors.New("id is empty")
	}

	err := s.repo.RepositoryDelete(id)
	if err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

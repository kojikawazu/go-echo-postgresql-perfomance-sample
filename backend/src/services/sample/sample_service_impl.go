package sample_service

import (
	sample_model "backend/src/models/sample"
	sample_repository "backend/src/repositories/sample"
)

// SampleService インターフェース
type SampleService interface {
	ServiceGetAllSamples() ([]sample_model.Sample, error)
	ServiceGetSampleByID(id string) (*sample_model.Sample, error)
	ServiceCreateSample(sample *sample_model.Sample) error
	ServiceUpdateSample(sample *sample_model.Sample) error
	ServiceDeleteSample(id string) error
}

// sampleService 構造体
type sampleService struct {
	repo sample_repository.SampleRepository
}

// NewSampleService 関数
func NewSampleService(repo sample_repository.SampleRepository) SampleService {
	return &sampleService{repo}
}

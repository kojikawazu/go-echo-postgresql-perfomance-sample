package sample_repository

import (
	sample_model "backend/src/models/sample"

	"gorm.io/gorm"
)

// SampleRepository インターフェース
type SampleRepository interface {
	RepositoryGetAll() ([]sample_model.Sample, error)
	RepositoryGetByID(id string) (*sample_model.Sample, error)
	RepositoryCreate(sample *sample_model.Sample) error
	RepositoryUpdate(sample *sample_model.Sample) error
	RepositoryDelete(id string) error
}

// sampleRepository 構造体
type sampleRepository struct {
	db *gorm.DB
}

// NewSampleRepository 関数
func NewSampleRepository(db *gorm.DB) SampleRepository {
	return &sampleRepository{db}
}

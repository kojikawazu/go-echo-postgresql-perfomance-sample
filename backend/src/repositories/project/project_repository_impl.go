package project_repository

import (
	project_model "backend/src/models/project"

	"gorm.io/gorm"
)

// ProjectRepository インターフェース
type ProjectRepository interface {
	RepositoryGetAll() ([]project_model.Project, error)
	RepositoryGetByID(id string) (*project_model.Project, error)
	RepositoryCreate(project *project_model.Project) error
	RepositoryUpdate(project *project_model.Project) error
	RepositoryDelete(id string) error
}

// projectRepository 構造体
type projectRepository struct {
	db *gorm.DB
}

// NewProjectRepository 関数
func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db}
}

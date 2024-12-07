package project_service

import (
	project_model "backend/src/models/project"
	project_repository "backend/src/repositories/project"
)

// ProjectService インターフェース
type ProjectService interface {
	ServiceGetAllProjects() ([]project_model.Project, error)
	ServiceGetProjectByID(id string) (*project_model.Project, error)
	ServiceCreateProject(project *project_model.Project) error
	ServiceUpdateProject(project *project_model.Project) error
	ServiceDeleteProject(id string) error
}

// projectService 構造体
type projectService struct {
	repo project_repository.ProjectRepository
}

// NewProjectService 関数
func NewProjectService(repo project_repository.ProjectRepository) ProjectService {
	return &projectService{repo}
}

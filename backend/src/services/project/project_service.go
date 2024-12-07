package project_service

import (
	project_model "backend/src/models/project"
	logging_utils "backend/src/utils/logging"
	"errors"
)

// 全権取得
func (s *projectService) ServiceGetAllProjects() ([]project_model.Project, error) {
	start := logging_utils.LogStart()

	projects, err := s.repo.RepositoryGetAll()
	if err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return projects, nil
}

// ID取得
func (s *projectService) ServiceGetProjectByID(id string) (*project_model.Project, error) {
	start := logging_utils.LogStart()

	// バリデーション
	if id == "" {
		logging_utils.LogError(start, errors.New("id is required"))
		return nil, errors.New("id is required")
	}

	project, err := s.repo.RepositoryGetByID(id)
	if err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogEnd(start)
	return project, nil
}

// 作成
func (s *projectService) ServiceCreateProject(project *project_model.Project) error {
	start := logging_utils.LogStart()

	// バリデーション
	if project.Name == "" {
		logging_utils.LogError(start, errors.New("name is required"))
		return errors.New("name is required")
	}
	if project.Description == "" {
		logging_utils.LogError(start, errors.New("description is required"))
		return errors.New("description is required")
	}

	err := s.repo.RepositoryCreate(project)
	if err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

// 更新
func (s *projectService) ServiceUpdateProject(project *project_model.Project) error {
	start := logging_utils.LogStart()

	// バリデーション
	if project.ID == "" {
		logging_utils.LogError(start, errors.New("id is required"))
		return errors.New("id is required")
	}
	if project.Name == "" {
		logging_utils.LogError(start, errors.New("name is required"))
		return errors.New("name is required")
	}
	if project.Description == "" {
		logging_utils.LogError(start, errors.New("description is required"))
		return errors.New("description is required")
	}

	err := s.repo.RepositoryUpdate(project)
	if err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

// 削除
func (s *projectService) ServiceDeleteProject(id string) error {
	start := logging_utils.LogStart()

	// バリデーション
	if id == "" {
		logging_utils.LogError(start, errors.New("id is required"))
		return errors.New("id is required")
	}

	err := s.repo.RepositoryDelete(id)
	if err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogEnd(start)
	return nil
}

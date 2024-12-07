package project_handler

import project_service "backend/src/services/project"

// プロジェクトハンドラ
type ProjectHandler struct {
	service project_service.ProjectService
}

// コンストラクタ
func NewProjectHandler(service project_service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service}
}

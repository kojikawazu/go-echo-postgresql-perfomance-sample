package project_handler

import (
	"net/http"

	project_model "backend/src/models/project"
	logging_utils "backend/src/utils/logging"

	"github.com/labstack/echo/v4"
)

// 全件取得
func (h *ProjectHandler) HandlerGetAllProjects(c echo.Context) error {
	start := logging_utils.LogStart()

	// 取得
	projects, err := h.service.ServiceGetAllProjects()
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	logging_utils.LogInfo("end projects:", projects)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, projects)
}

// 1件取得
func (h *ProjectHandler) HandlerGetProjectByID(c echo.Context) error {
	start := logging_utils.LogStart()

	id := c.Param("id")

	// 取得
	project, err := h.service.ServiceGetProjectByID(id)
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	logging_utils.LogInfo("end project:", project)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, project)
}

// 作成
func (h *ProjectHandler) HandlerCreateProject(c echo.Context) error {
	start := logging_utils.LogStart()

	var project project_model.Project
	if err := c.Bind(&project); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// 作成
	if err := h.service.ServiceCreateProject(&project); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	logging_utils.LogInfo("end project:", project)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, project)
}

// 更新
func (h *ProjectHandler) HandlerUpdateProject(c echo.Context) error {
	start := logging_utils.LogStart()

	id := c.Param("id")
	var project project_model.Project
	if err := c.Bind(&project); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// 更新
	project.ID = id
	if err := h.service.ServiceUpdateProject(&project); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	logging_utils.LogInfo("end project:", project)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, project)
}

// 削除
func (h *ProjectHandler) HandlerDeleteProject(c echo.Context) error {
	start := logging_utils.LogStart()

	id := c.Param("id")

	// 削除
	err := h.service.ServiceDeleteProject(id)
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	logging_utils.LogInfo("end project:", id)
	logging_utils.LogEnd(start)
	return c.NoContent(http.StatusNoContent)
}

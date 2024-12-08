package task_handler

import (
	"net/http"

	task_model "backend/src/models/task"
	logging_utils "backend/src/utils/logging"

	"github.com/labstack/echo/v4"
)

// 全件取得
func (h *TaskHandler) HandlerGetAllTasks(c echo.Context) error {
	start := logging_utils.LogStart()

	// 取得
	tasks, err := h.service.ServiceGetAllTasks()
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//logging_utils.LogInfo("end tasks:", tasks)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, tasks)
}

// 1件取得
func (h *TaskHandler) HandlerGetTaskByID(c echo.Context) error {
	start := logging_utils.LogStart()

	id := c.Param("id")

	// 取得
	task, err := h.service.ServiceGetTaskByID(id)
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//logging_utils.LogInfo("end task:", task)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, task)
}

// 作成
func (h *TaskHandler) HandlerCreateTask(c echo.Context) error {
	start := logging_utils.LogStart()

	var task task_model.Task
	if err := c.Bind(&task); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// 作成
	if err := h.service.ServiceCreateTask(&task); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//logging_utils.LogInfo("end task:", task)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, task)
}

// 更新
func (h *TaskHandler) HandlerUpdateTask(c echo.Context) error {
	start := logging_utils.LogStart()

	id := c.Param("id")
	var task task_model.Task
	if err := c.Bind(&task); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// 更新
	task.ID = id
	if err := h.service.ServiceUpdateTask(&task); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//logging_utils.LogInfo("end task:", task)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, task)
}

// 削除
func (h *TaskHandler) HandlerDeleteTask(c echo.Context) error {
	start := logging_utils.LogStart()

	id := c.Param("id")

	// 削除
	err := h.service.ServiceDeleteTask(id)
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//logging_utils.LogInfo("end task:", id)
	logging_utils.LogEnd(start)
	return c.NoContent(http.StatusNoContent)
}

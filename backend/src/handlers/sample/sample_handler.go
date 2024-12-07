package sample_handler

import (
	"net/http"

	sample_model "backend/src/models/sample"
	logging_utils "backend/src/utils/logging"

	"github.com/labstack/echo/v4"
)

// 全件取得
func (h *SampleHandler) HandleGetAllSamples(c echo.Context) error {
	start := logging_utils.LogStart()

	samples, err := h.service.ServiceGetAllSamples()
	//エラーのタイプで分岐
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	logging_utils.LogInfo("HandleGetAllSamples end samples:", samples)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, samples)
}

// 1件取得
func (h *SampleHandler) HandleGetSampleByID(c echo.Context) error {
	start := logging_utils.LogStart()

	id := c.Param("id")
	sample, err := h.service.ServiceGetSampleByID(id)
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Sample not found"})
	}

	logging_utils.LogInfo("HandleGetSampleByID end sample:", sample)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, sample)
}

// 作成
func (h *SampleHandler) HandleCreateSample(c echo.Context) error {
	start := logging_utils.LogStart()

	var sample sample_model.Sample
	if err := c.Bind(&sample); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.service.ServiceCreateSample(&sample); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	logging_utils.LogInfo("HandleCreateSample end sample:", sample)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusCreated, sample)
}

// 更新
func (h *SampleHandler) HandleUpdateSample(c echo.Context) error {
	start := logging_utils.LogStart()

	id := c.Param("id")
	var sample sample_model.Sample
	if err := c.Bind(&sample); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	sample.ID = id
	if err := h.service.ServiceUpdateSample(&sample); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	logging_utils.LogInfo("HandleUpdateSample end sample:", sample)
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, sample)
}

// 削除
func (h *SampleHandler) HandleDeleteSample(c echo.Context) error {
	start := logging_utils.LogStart()

	id := c.Param("id")
	if err := h.service.ServiceDeleteSample(id); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	logging_utils.LogInfo("HandleDeleteSample end")
	logging_utils.LogEnd(start)
	return c.NoContent(http.StatusNoContent)
}

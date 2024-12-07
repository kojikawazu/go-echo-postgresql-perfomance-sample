package routes

import (
	sample_handler "backend/src/handlers/sample"

	"github.com/labstack/echo/v4"
)

func SetupSampleRoutes(e *echo.Echo, sampleHandler *sample_handler.SampleHandler) {
	sampleGroup := e.Group("/samples")
	sampleGroup.GET("", sampleHandler.HandleGetAllSamples)
	sampleGroup.GET("/:id", sampleHandler.HandleGetSampleByID)
	sampleGroup.POST("", sampleHandler.HandleCreateSample)
	sampleGroup.PUT("/:id", sampleHandler.HandleUpdateSample)
	sampleGroup.DELETE("/:id", sampleHandler.HandleDeleteSample)
}

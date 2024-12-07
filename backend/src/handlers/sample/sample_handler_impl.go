package sample_handler

import sample_service "backend/src/services/sample"

type SampleHandler struct {
	service sample_service.SampleService
}

func NewSampleHandler(service sample_service.SampleService) *SampleHandler {
	return &SampleHandler{service}
}

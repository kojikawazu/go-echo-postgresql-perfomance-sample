package sample_handler

import sample_service "backend/src/services/sample"

// サンプルハンドラ
type SampleHandler struct {
	service sample_service.SampleService
}

// コンストラクタ
func NewSampleHandler(service sample_service.SampleService) *SampleHandler {
	return &SampleHandler{service}
}

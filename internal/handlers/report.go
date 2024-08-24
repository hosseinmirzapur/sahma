package handlers

import "github.com/gin-gonic/gin"

type reportHandler struct{}

func ReportHandler() *reportHandler {
	return &reportHandler{}
}

func (h *reportHandler) UsersReport(c *gin.Context) {}

func (h *reportHandler) CreateExcelUserReport(c *gin.Context) {}

func (h *reportHandler) TotalUploadedFiles(c *gin.Context) {}

func (h *reportHandler) TotalUploadedFilesByType(c *gin.Context) {}

func (h *reportHandler) TotalTranscribedFiles(c *gin.Context) {}

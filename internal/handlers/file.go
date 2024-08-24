package handlers

import "github.com/gin-gonic/gin"

type fileHandler struct{}

func FileHandler() *fileHandler {
	return &fileHandler{}
}

func (h *fileHandler) Show(c *gin.Context) {}

func (h *fileHandler) AddDescription(c *gin.Context) {}

func (h *fileHandler) Transcribe(c *gin.Context) {}

func (h *fileHandler) DownloadOriginalFile(c *gin.Context) {}

func (h *fileHandler) DownloadSearchableFile(c *gin.Context) {}

func (h *fileHandler) DownloadWordFile(c *gin.Context) {}

func (h *fileHandler) Rename(c *gin.Context) {}

func (h *fileHandler) PrintOriginalFile(c *gin.Context) {}

func (h *fileHandler) Upload(c *gin.Context) {}

func (h *fileHandler) UploadRoot(c *gin.Context) {}

package handlers

import "github.com/gin-gonic/gin"

type dashboardHandler struct{}

func DashboardHandler() *dashboardHandler {
	return &dashboardHandler{}
}

func (h *dashboardHandler) Index(c *gin.Context) {}

func (h *dashboardHandler) Copy(c *gin.Context) {}

func (h *dashboardHandler) Move(c *gin.Context) {}

func (h *dashboardHandler) PermanentDelete(c *gin.Context) {}

func (h *dashboardHandler) TrashList(c *gin.Context) {}

func (h *dashboardHandler) TrashAction(c *gin.Context) {}

func (h *dashboardHandler) TrashRetrieve(c *gin.Context) {}

func (h *dashboardHandler) ArchiveList(c *gin.Context) {}

func (h *dashboardHandler) ArchiveAction(c *gin.Context) {}

func (h *dashboardHandler) ArchiveRetrieve(c *gin.Context) {}

func (h *dashboardHandler) CreateZip(c *gin.Context) {}

func (h *dashboardHandler) SearchForm(c *gin.Context) {}

func (h *dashboardHandler) SearchAction(c *gin.Context) {}

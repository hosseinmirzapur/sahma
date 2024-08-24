package handlers

import "github.com/gin-gonic/gin"

type folderHandler struct{}

func FolderHandler() *folderHandler {
	return &folderHandler{}
}

func (h *folderHandler) Show(c *gin.Context) {}

func (h *folderHandler) CreateRoot(c *gin.Context) {}

func (h *folderHandler) Create(c *gin.Context) {}

func (h *folderHandler) Rename(c *gin.Context) {}

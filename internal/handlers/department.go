package handlers

import "github.com/gin-gonic/gin"

type departmentHandler struct{}

func DepartmentHandler() *departmentHandler {
	return &departmentHandler{}
}

func (h *departmentHandler) Index(c *gin.Context) {}

func (h *departmentHandler) Create(c *gin.Context) {}

func (h *departmentHandler) Edit(c *gin.Context) {}

func (h *departmentHandler) Delete(c *gin.Context) {}

package handlers

import "github.com/gin-gonic/gin"

type userHandler struct{}

func UserHandler() *userHandler {
	return &userHandler{}
}

func (h *userHandler) Index(c *gin.Context) {}

func (h *userHandler) UserInfo(c *gin.Context) {}

func (h *userHandler) Create(c *gin.Context) {}

func (h *userHandler) Block(c *gin.Context) {}

func (h *userHandler) Edit(c *gin.Context) {}

func (h *userHandler) Search(c *gin.Context) {}

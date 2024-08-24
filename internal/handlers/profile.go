package handlers

import "github.com/gin-gonic/gin"

type profileHandler struct{}

func ProfileHandler() *profileHandler {
	return &profileHandler{}
}

func (h *profileHandler) Show(c *gin.Context) {}

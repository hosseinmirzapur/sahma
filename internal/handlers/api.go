package handlers

import "github.com/gin-gonic/gin"

type apiHandler struct{}

func APIHandler() *apiHandler {
	return &apiHandler{}
}

func (h *apiHandler) ListUsers(c *gin.Context) {}

func (h *apiHandler) ListLetters(c *gin.Context) {}

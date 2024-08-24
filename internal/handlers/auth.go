package handlers

import "github.com/gin-gonic/gin"

type authHandler struct{}

func AuthHandler() *authHandler {
	return &authHandler{}
}

func (h *authHandler) Index(c *gin.Context) {}

func (h *authHandler) LoginPage(c *gin.Context) {}

func (h *authHandler) LoginAction(c *gin.Context) {}

func (h *authHandler) Logout(c *gin.Context) {}

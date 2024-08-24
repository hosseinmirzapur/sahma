package handlers

import "github.com/gin-gonic/gin"

type notificationHandler struct{}

func NotificationHandler() *notificationHandler {
	return &notificationHandler{}
}

func (h *notificationHandler) Index(c *gin.Context) {}

func (h *notificationHandler) CreateAction(c *gin.Context) {}

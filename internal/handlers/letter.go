package handlers

import "github.com/gin-gonic/gin"

type letterHandler struct{}

func LetterHandler() *letterHandler {
	return &letterHandler{}
}

func (h *letterHandler) Inbox(c *gin.Context) {}

func (h *letterHandler) GetDraftedLetters(c *gin.Context) {}

func (h *letterHandler) GetSubmittedLetters(c *gin.Context) {}

func (h *letterHandler) GetDeletedLetters(c *gin.Context) {}

func (h *letterHandler) GetArchivedLetters(c *gin.Context) {}

func (h *letterHandler) SubmitForm(c *gin.Context) {}

func (h *letterHandler) SubmitAction(c *gin.Context) {}

func (h *letterHandler) Show(c *gin.Context) {}

func (h *letterHandler) SignAction(c *gin.Context) {}

func (h *letterHandler) ReferAction(c *gin.Context) {}

func (h *letterHandler) ReplyAction(c *gin.Context) {}

func (h *letterHandler) DownloadAttachment(c *gin.Context) {}

func (h *letterHandler) DraftAction(c *gin.Context) {}

func (h *letterHandler) ShowDrafted(c *gin.Context) {}

func (h *letterHandler) SubmitDrafted(c *gin.Context) {}

func (h *letterHandler) Archive(c *gin.Context) {}

func (h *letterHandler) TempDelete(c *gin.Context) {}

func (h *letterHandler) SubmitNotification(c *gin.Context) {}

package models

import (
	"fmt"
	"time"
)

type Letter struct {
	ID                  uint
	UserID              uint
	Subject             string
	Text                *string
	Status              string
	Description         *string
	Meta                *string
	Priority            string
	SubmittedAt         *time.Time
	DueDate             *time.Time
	Category            string
	CreatedAt           time.Time
	UpdatedAt           *time.Time
	LetterReferenceType *string
	LetterReferenceID   *uint
	Activities          []*Activity         `gorm:"polymorphic:Activity"`
	Attachments         []*LetterAttachment `gorm:"polymorphic:Attachable"`
	LetterInboxes       []*LetterInbox
	LetterReplies       []*LetterReply
	LetterSigns         []*LetterSign
	Notifications       []*Notification
	User                User
}

// Priorities
var (
	LETTER_PRIORITY_NORMAL      = "NORMAL"
	LETTER_PRIORITY_IMMEDIATELY = "IMMEDIATELY"
	LETTER_PRIORITY_INSTANT     = "INSTANT"
)

// Categories
var (
	LETTER_CATEGORY_NORMAL       = "NORMAL"
	LETTER_CATEGORY_SECRET       = "SECRET"
	LETTER_CATEGORY_CONFIDENTIAL = "CONFIDENTIAL"
)

// Different types of Status
var (
	LETTER_STATUS_SENT     = "SENT"
	LETTER_STATUS_RECEIVED = "RECEIVED"
	LETTER_STATUS_REPLIED  = "REPLIED"
	LETTER_STATUS_ACHIEVED = "ACHIEVED"
	LETTER_STATUS_DELETED  = "DELETED"
	LETTER_STATUS_DRAFT    = "DRAFT"
)

func (l *Letter) GetLetterStatus(user User) string {
	status := ""
	if l.Status == LETTER_STATUS_SENT {
		if l.UserID == user.ID {
			status = LETTER_STATUS_SENT
		} else {
			status = LETTER_STATUS_RECEIVED
		}
	} else {
		status = l.Status
	}

	return status
}

func (l *Letter) GetPriorityLetterInPersian() (string, error) {
	switch l.Priority {
	case LETTER_PRIORITY_NORMAL:
		return "عادی", nil
	case LETTER_PRIORITY_IMMEDIATELY:
		return "فوری", nil
	case LETTER_PRIORITY_INSTANT:
		return "آنی", nil
	default:
		return "", fmt.Errorf("%v", map[string]interface{}{
			"message": "unsupported priority letter",
		})
	}
}

func (l *Letter) GetCategoryLetterInPersian() (string, error) {
	switch l.Category {
	case LETTER_CATEGORY_SECRET:
		return "سری", nil
	case LETTER_CATEGORY_CONFIDENTIAL:
		return "مطمئن", nil
	case LETTER_CATEGORY_NORMAL:
		return "عادی", nil
	default:
		return "", fmt.Errorf("%v", map[string]interface{}{
			"message": "unsupported category letter",
		})
	}
}

func GetAllLetterPriorities() []string {
	return []string{
		LETTER_PRIORITY_INSTANT, LETTER_PRIORITY_IMMEDIATELY, LETTER_PRIORITY_NORMAL,
	}
}

func GetmimeTypes() []string {
	return []string{"jpeg", "jpg", "png", "pdf"}
}

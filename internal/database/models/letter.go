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
	PRIORITY_NORMAL      = "NORMAL"
	PRIORITY_IMMEDIATELY = "IMMEDIATELY"
	PRIORITY_INSTANT     = "INSTANT"
)

// Categories
var (
	CATEGORY_NORMAL       = "NORMAL"
	CATEGORY_SECRET       = "SECRET"
	CATEGORY_CONFIDENTIAL = "CONFIDENTIAL"
)

// Different types of Status
var (
	STATUS_SENT     = "SENT"
	STATUS_RECEIVED = "RECEIVED"
	STATUS_REPLIED  = "REPLIED"
	STATUS_ACHIEVED = "ACHIEVED"
	STATUS_DELETED  = "DELETED"
	STATUS_DRAFT    = "DRAFT"
)

func (l *Letter) GetLetterStatus(user User) string {
	status := ""
	if l.Status == STATUS_SENT {
		if l.UserID == user.ID {
			status = STATUS_SENT
		} else {
			status = STATUS_RECEIVED
		}
	} else {
		status = l.Status
	}

	return status
}

func (l *Letter) GetPriorityLetterInPersian() (string, error) {
	switch l.Priority {
	case PRIORITY_NORMAL:
		return "عادی", nil
	case PRIORITY_IMMEDIATELY:
		return "فوری", nil
	case PRIORITY_INSTANT:
		return "آنی", nil
	default:
		return "", fmt.Errorf("%v", map[string]interface{}{
			"message": "unsupported priority letter",
		})
	}
}

func (l *Letter) GetCategoryLetterInPersian() (string, error) {
	switch l.Category {
	case CATEGORY_SECRET:
		return "سری", nil
	case CATEGORY_CONFIDENTIAL:
		return "مطمئن", nil
	case CATEGORY_NORMAL:
		return "عادی", nil
	default:
		return "", fmt.Errorf("unsupported category letter")
	}
}

func GetAllLetterPriorities() []string {
	return []string{
		PRIORITY_INSTANT, PRIORITY_IMMEDIATELY, PRIORITY_NORMAL,
	}
}

func GetmimeTypes() []string {
	return []string{"jpeg", "jpg", "png", "pdf"}
}

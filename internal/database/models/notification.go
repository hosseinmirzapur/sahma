package models

import (
	"fmt"
	"time"
)

type Notification struct {
	ID          uint
	UserID      uint
	LetterID    *uint
	Subject     *string
	Description *string
	Priority    string
	Meta        *string
	RemindAt    *time.Time
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	Letter      *Letter
	User        User
}

// Priorities
var (
	NOTIFICATION_PRIORITY_NORMAL      = ""
	NOTIFICATION_PRIORITY_IMMEDIATELY = ""
	NOTIFICATION_PRIORITY_INSTANT     = ""
)

func (notif *Notification) GetPriorityNotification() (string, error) {
	switch notif.Priority {
	case NOTIFICATION_PRIORITY_NORMAL:
		return "عادی", nil
	case NOTIFICATION_PRIORITY_IMMEDIATELY:
		return "فوری", nil
	case NOTIFICATION_PRIORITY_INSTANT:
		return "آنی", nil
	default:
		return "", fmt.Errorf("%v", map[string]interface{}{
			"message": "unsupported priority notification",
		})
	}
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID           uint
	UserID       uint
	Status       string
	Description  *string
	ActivityType ActivityType
	ActivityID   string
	Meta         *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         User
}

// Activity types
type ActivityType string

const (
	TYPE_CREATE        ActivityType = "CREATE"
	TYPE_PRINT         ActivityType = "PRINT"
	TYPE_DESCRIPTION   ActivityType = "DESCRIPTION"
	TYPE_UPLOAD        ActivityType = "UPLOAD"
	TYPE_DELETE        ActivityType = "DELETE"
	TYPE_RENAME        ActivityType = "RENAME"
	TYPE_COPY          ActivityType = "COPY"
	TYPE_EDIT          ActivityType = "EDIT"
	TYPE_TRANSCRIPTION ActivityType = "TRANSCRIPTION"
	TYPE_LOGIN         ActivityType = "LOGIN"
	TYPE_LOGOUT        ActivityType = "LOGOUT"
	TYPE_ARCHIVE       ActivityType = "ARCHIVE"
	TYPE_RETRIEVAL     ActivityType = "RETRIEVAL"
	TYPE_MOVE          ActivityType = "MOVE"
	TYPE_DOWNLOAD      ActivityType = "DOWNLOAD"
)

func (a *Activity) ForPeriod(query *gorm.DB, start, end string) *gorm.DB {
	return query.Where("created_at BETWEEN ? AND ?", start, end)
}

func (a *Activity) Logins(query *gorm.DB) *gorm.DB {
	return query.Where("status", TYPE_LOGIN)
}

func (a *Activity) Logouts(query *gorm.DB) *gorm.DB {
	return query.Where("status", TYPE_LOGOUT)
}

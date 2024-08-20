package models

import "time"

type LetterInbox struct {
	ID               uint
	LetterID         *uint
	UserID           uint
	ReadStatus       string
	IsRefer          *string
	ReferredBy       uint
	ReferDescription *string
	DueDate          *string
	Meta             *string
	CreatedAt        time.Time
	UpdatedAt        *time.Time
	Letter           *Letter
	ReferrerUser     *User
	User             *User
}

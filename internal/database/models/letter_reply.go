package models

import "time"

type LetterReply struct {
	ID          uint
	LetterID    *uint
	UserID      uint
	RecipientID *uint
	Text        string
	Meta        *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	User        User
	Letter      *Letter
	Attachments []*LetterAttachment `gorm:"polymorphic:Attachable"`
}

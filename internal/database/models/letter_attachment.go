package models

import "time"

type LetterAttachment struct {
	ID             uint
	Type           string
	FileLocation   string
	AttachableType string
	AttachableID   uint
	Meta           *string
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	Letter         *Letter
}

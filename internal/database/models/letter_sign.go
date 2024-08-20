package models

import "time"

type LetterSign struct {
	ID        uint
	LetterID  *uint
	UserID    *uint
	SignedAt  *time.Time
	CreatedAt time.Time
	UpdatedAt *time.Time
	Letter    *Letter
	User      User
}

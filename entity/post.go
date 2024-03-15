package entity

import "github.com/google/uuid"

type Post struct {
	ID          uint  `json:"id" gorm:"primary_key;autoIncrement"`
	Title       string `json:"title" gorm:"type:varchar(255);not null;"`
	Description string `json:"description" gorm:"type:text;not null;"`
	User        User   `json:"-"`
	UserID      uuid.UUID
}

// uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`

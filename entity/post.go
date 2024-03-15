package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uint   `json:"id" gorm:"primary_key;autoIncrement"`
	Title       string `json:"title" gorm:"type:varchar(255);not null;"`
	Description string `json:"description" gorm:"type:text;not null;"`
	User        User   `json:"-"`
	UserID      uuid.UUID
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
// `json:"id" gorm:"type:varchar(36);primary_key;"`

// `json:"id" gorm:"primary_key;autoIncrement"`

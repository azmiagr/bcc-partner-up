package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null;"`
	RoleID    int
	Role      Role      `json:"roles,omitempty" gorm:"foreignKey:RoleID;references:ID"`
	Alamat    string    `json:"alamat" gorm:"type:varchar(255);"`
	Minat     string    `json:"minat" gorm:"type:varchar(255);"`
	Skill     string    `json:"skill" gorm:"type:varchar(255);"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Post      []Post
}

// `json:"role" gorm:"foreinKey:ID; references:roles; not null;"`

// lama
// `json:"roles,omitempty" gorm:"foreignKey:RoleID;references:ID"`

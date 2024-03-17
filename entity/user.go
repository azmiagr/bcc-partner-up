package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name     string    `json:"name" gorm:"type:varchar(255);not null;"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string    `json:"password" gorm:"type:varchar(255);not null;"`
	RoleID   uint
	// Role       Role      `json:"roles,omitempty" gorm:"foreignKey:RoleID;references:ID"`
	Alamat     string    `json:"alamat" gorm:"type:varchar(255);"`
	Minat      []*Minat  `json:"minat" gorm:"many2many:user_minat;"`
	Skill      []*Skill  `json:"skill" gorm:"many2many:user_skill;"`
	PhotoLink  string    `json:"photoLink" gorm:"type:varchar(200);"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Post       []Post
	UniID      uint
	DistrictID uint
}

// `json:"role" gorm:"foreinKey:ID; references:roles; not null;"`

// lama
// `json:"roles,omitempty" gorm:"foreignKey:RoleID;references:ID"`

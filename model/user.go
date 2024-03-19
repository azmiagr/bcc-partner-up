package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UserRegister struct {
	ID       uuid.UUID `json:"-"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8"`
	RoleID   int

	// Uni      int       `json:"uni" binding:"required"`
	// District int       `json:"district" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserParam struct {
	ID       uuid.UUID `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
}

type GetUserByNameResponse struct {
	Name  string
	Uni   uint
	Minat []uint `json:"minat_id"`
	Skill []uint `json:"skill_id"`
}

type UploadPhoto struct {
	Photo *multipart.FileHeader `form:"photo"`
}

type UpdateProfile struct {
	Name     string `json:"name"`
	Uni      uint   `json:"uni_id"`
	District uint   `json:"district_id"`
	Minat    []uint `json:"minat_id" binding:"lte=5"`
	Skill    []uint `json:"skill_id" binding:"lte=5"`
}

type UpdateProfileResponse struct {
	Name     string `json:"name"`
	Uni      uint   `json:"uni_id"`
	District uint   `json:"district_id"`
	Minat    []uint `json:"minat_id"`
	Skill    []uint `json:"skill_id"`
}

// type UserRegister struct {
// 	ID       uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
// 	Email    string    `json:"email" gorm:"type:varchar(255);not null;unique"`
// 	Password string    `json:"password" gorm:"type:varchar(255);not null;"`
// }

/*
Nama lengkap
Alamat: Kota, Kabupaten dan kecamatan
Asal universitas (tahun mulai tahun akhir)
Minat
Skill

Yang dibutuhin pas register
Email
Password
(verifikasi)


*/

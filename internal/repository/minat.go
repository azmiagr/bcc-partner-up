package repository

import (
	"intern-bcc/entity"

	"gorm.io/gorm"
)

type IMinatRepository interface {
	GetAllMinat() ([]entity.Minat, error)
}

type MinatRepository struct {
	db *gorm.DB
}

func NewUMinatRepository(db *gorm.DB) IMinatRepository {
	return &MinatRepository{db}
}

func (mr *MinatRepository) GetAllMinat() ([]entity.Minat, error) {
	var minat []entity.Minat
	if err := mr.db.Debug().Find(&minat).Error; err != nil {
		return nil, err
	}

	return minat, nil
}

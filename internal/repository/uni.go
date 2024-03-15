package repository

import (
	"intern-bcc/entity"

	"gorm.io/gorm"
)

type IUniRepository interface {
	GetAllUni() ([]entity.Uni, error)
}

type UniRepo struct {
	db *gorm.DB
}

func NewUniRrepository(db *gorm.DB) IUniRepository {
	return &UniRepo{db}
}

func (br *UniRepo) GetAllUni() ([]entity.Uni, error) {
	var uni []entity.Uni
	if err := br.db.Debug().Find(&uni).Error; err != nil {
		return nil, err
	}

	return uni, nil
}

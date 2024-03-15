package repository

import (
	"intern-bcc/entity"

	"gorm.io/gorm"
)

type IDistrictRepository interface {
	GetAllDistrict() ([]entity.District, error)
}

type DistrictRepo struct {
	db *gorm.DB
}

func NewDistrictRepo(db *gorm.DB) IDistrictRepository {
	return &DistrictRepo{db}
}

func (dr *DistrictRepo) GetAllDistrict() ([]entity.District, error) {
	var district []entity.District
	if err := dr.db.Debug().Find(&district).Error; err != nil {
		return nil, err
	}

	return district, nil
}

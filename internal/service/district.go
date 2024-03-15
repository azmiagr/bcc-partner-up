package service

import (
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
)

type IDistrictService interface {
	GetAllDistrict() ([]entity.District, error)
}

type DistrictService struct {
	DistrictRepo repository.IDistrictRepository
}

func NewDistrictService(districtService repository.IDistrictRepository) IDistrictService {
	return &DistrictService{districtService}
}

func (ds *DistrictService) GetAllDistrict() ([]entity.District, error) {
	district, err := ds.DistrictRepo.GetAllDistrict()
	if err != nil {
		return nil, err
	}

	return district, nil
}

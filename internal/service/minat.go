package service

import (
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
)

type IMinatService interface {
	GetAllMinat() ([]entity.Minat, error)
}

type MinatService struct {
	MinatRepo repository.IMinatRepository
}

func NewMinatService(minatRepo repository.IMinatRepository) IMinatService {
	return &MinatService{minatRepo}
}

func (ms *MinatService) GetAllMinat() ([]entity.Minat, error) {
	minat, err := ms.MinatRepo.GetAllMinat()
	if err != nil {
		return nil, err
	}

	return minat, nil
}

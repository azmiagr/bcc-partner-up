package service

import (
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
)

type IUniService interface {
	GetAllUni() ([]entity.Uni, error)
}

type UniService struct {
	UniRepo repository.IUniRepository
}

func NewUniService(uniRepo repository.IUniRepository) IUniService {
	return &UniService{uniRepo}
}

func (us *UniService) GetAllUni() ([]entity.Uni, error) {
	uni, err := us.UniRepo.GetAllUni()
	if err != nil {
		return nil, err
	}

	return uni, nil
}

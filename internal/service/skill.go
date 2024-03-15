package service

import (
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
)

type ISkillService interface {
	GetAllSkill() ([]entity.Skill, error)
}

type SkillService struct {
	SkillRepo repository.ISkillRepository
}

func NewSkillService(skillRepo repository.ISkillRepository) ISkillService {
	return &SkillService{skillRepo}
}

func (sr *SkillService) GetAllSkill() ([]entity.Skill, error) {
	skill, err := sr.SkillRepo.GetAllSkill()
	if err != nil {
		return nil, err
	}

	return skill, nil
}

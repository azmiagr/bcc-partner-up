package repository

import (
	"intern-bcc/entity"

	"gorm.io/gorm"
)

type ISkillRepository interface {
	GetAllSkill() ([]entity.Skill, error)
}

type SkillRepo struct {
	db *gorm.DB
}

func NewSkillRrepository(db *gorm.DB) ISkillRepository {
	return &SkillRepo{db}
}

func (sr *SkillRepo) GetAllSkill() ([]entity.Skill, error) {
	var skill []entity.Skill
	if err := sr.db.Debug().Find(&skill).Error; err != nil {
		return nil, err
	}

	return skill, nil
}

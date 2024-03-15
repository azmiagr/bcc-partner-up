package repository

import "gorm.io/gorm"

type Repository struct {
	User     IUserRepository
	Post     IPostRepository
	Uni      IUniRepository
	District IDistrictRepository
	Minat    IMinatRepository
	Skill    ISkillRepository
}

func NewRepository(db *gorm.DB) *Repository {
	UserRepo := NewUserRepository(db)
	PostRepo := NewPostRepository(db)
	UniRepo := NewUniRrepository(db)
	DistrictRepo := NewDistrictRepo(db)
	MinatRepo := NewUMinatRepository(db)
	SkillRepo := NewSkillRrepository(db)
	return &Repository{
		User:     UserRepo,
		Post:     PostRepo,
		Uni:      UniRepo,
		District: DistrictRepo,
		Minat:    MinatRepo,
		Skill:    SkillRepo,
	}
}

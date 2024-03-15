package service

import (
	"intern-bcc/internal/repository"
	"intern-bcc/pkg/bcrypt"
	"intern-bcc/pkg/jwt"
	"intern-bcc/pkg/supabase"
	// "intern-bcc/model"
)

type Service struct {
	User     IUserService
	Post     IPostService
	Uni      IUniService
	District IDistrictService
	Minat    IMinatService
	Skill    ISkillService
}

func NewService(r *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, supabase supabase.Interface) *Service {
	return &Service{
		User:     NewUserService(r.User, bcrypt, jwtAuth, supabase),
		Post:     NewPostService(r.Post),
		Uni:      NewUniService(r.Uni),
		District: NewDistrictService(r.District),
		Minat:    NewMinatService(r.Minat),
		Skill:    NewSkillService(r.Skill),
	}
}

// func (u *UserService) Login(param model.UserLogin) (model.UserLoginResponse, error) {

// }

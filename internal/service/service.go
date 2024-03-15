package service

import (
	"intern-bcc/internal/repository"
	"intern-bcc/pkg/bcrypt"
	"intern-bcc/pkg/jwt"
	// "intern-bcc/model"
)

type Service struct {
	User     IUserService
	Post     IPostService
	Uni      IUniService
	District IDistrictService
}

func NewService(r *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) *Service {
	return &Service{
		User:     NewUserService(r.User, bcrypt, jwtAuth),
		Post:     NewPostService(r.Post),
		Uni:      NewUniService(r.Uni),
		District: NewDistrictService(r.District),
	}
}

// func (u *UserService) Login(param model.UserLogin) (model.UserLoginResponse, error) {

// }

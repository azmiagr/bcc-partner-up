package service

import (
	"errors"
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
	"intern-bcc/model"
	"intern-bcc/pkg/bcrypt"
	"intern-bcc/pkg/jwt"

	"github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Register(data model.UserRegister) error
	Login(param model.UserLogin) (model.UserLoginResponse, error)
	GetUser(param model.UserParam) (entity.User, error)
	GetUserByName(name string) (*entity.User, error)
}

type UserService struct {
	UserRepo repository.IUserRepository
	bcrypt   bcrypt.Interface
	jwtAuth  jwt.Interface
}

func NewUserService(r repository.IUserRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) IUserService {
	return &UserService{
		UserRepo: r,
		bcrypt:   bcrypt,
		jwtAuth:  jwtAuth,
	}
}

func (u *UserService) Register(data model.UserRegister) error {
	hash, err := u.bcrypt.GenerateFromPasswordstring(data.Password)

	if err != nil {
		return err
	}

	id, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	user := &entity.User{
		ID:       id,
		Email:    data.Email,
		Password: hash,
		RoleID:   2, // nanti ganti !!!!!!!
	}

	err = u.UserRepo.CreateUser(user)

	if err != nil {
		return err
	}

	return nil

}

func (u *UserService) Login(param model.UserLogin) (model.UserLoginResponse, error) {
	var result = model.UserLoginResponse{}
	user, err := u.UserRepo.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return result, err
	}

	err = u.bcrypt.CompareAndHashPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := u.jwtAuth.CreateJWTToken(user.ID)
	if err != nil {
		return result, errors.New("failed create jwt")
	}

	result.Token = token

	return result, nil

}

func (u *UserService) GetUser(param model.UserParam) (entity.User, error) {
	return u.UserRepo.GetUser(param)
}

func (u *UserService) GetUserByName(name string) (*entity.User, error) {
	user, err := u.UserRepo.GetUserByName(name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

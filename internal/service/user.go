package service

import (
	"errors"
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
	"intern-bcc/model"
	"intern-bcc/pkg/bcrypt"
	"intern-bcc/pkg/jwt"
	"intern-bcc/pkg/supabase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Register(data *model.UserRegister) (*entity.User, error)
	Login(param model.UserLogin) (model.UserLoginResponse, error)
	GetUser(param model.UserParam) (entity.User, error)
	GetUserByName(name string) (*entity.User, error)
	UploadPhoto(ctx *gin.Context, param model.UploadPhoto) error
	UpdateProfile(id string, profileReq *model.UpdateProfile) (*entity.User, error)
	GetUsersByFilter(uniID uint, minatID []uint, skillID []uint) ([]model.UserFilter, error)
}

type UserService struct {
	UserRepo repository.IUserRepository
	bcrypt   bcrypt.Interface
	jwtAuth  jwt.Interface
	supabase supabase.Interface
}

func NewUserService(r repository.IUserRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, supabase supabase.Interface) IUserService {
	return &UserService{
		UserRepo: r,
		bcrypt:   bcrypt,
		jwtAuth:  jwtAuth,
		supabase: supabase,
	}
}

func (u *UserService) Register(data *model.UserRegister) (*entity.User, error) {
	hash, err := u.bcrypt.GenerateFromPasswordstring(data.Password)

	if err != nil {
		return nil, err
	}

	id, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:         id,
		Email:      data.Email,
		Password:   hash,
		RoleID:     2,
		UniID:      1,
		DistrictID: 1,
	}

	user, err = u.UserRepo.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return user, err

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

func (u *UserService) UploadPhoto(ctx *gin.Context, param model.UploadPhoto) error {
	user, err := u.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		return err
	}

	if user.PhotoLink != "" {
		err := u.supabase.Delete(user.PhotoLink)
		if err != nil {
			return err
		}
	}

	link, err := u.supabase.Upload(param.Photo)
	if err != nil {
		return err
	}

	err = u.UserRepo.UpdateUserPhoto(entity.User{
		PhotoLink: link,
	}, model.UserParam{
		ID: user.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserService) UpdateProfile(id string, profileReq *model.UpdateProfile) (*entity.User, error) {
	user, err := ur.UserRepo.UpdateProfile(id, profileReq)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetUsersByFilter(uniID uint, minatID []uint, skillID []uint) ([]model.UserFilter, error) {
	users, err := u.UserRepo.GetUsersByFilter(uniID, minatID, skillID)
	if err != nil {
		return nil, err
	}

	var filter []model.UserFilter
	for _, user := range users {
		var minatID []uint
		for _, minat := range user.Minat {
			minatID = append(minatID, minat.ID)
		}

		var skillID []uint
		for _, skill := range user.Skill {
			skillID = append(skillID, skill.ID)
		}

		res := model.UserFilter{
			Name:  user.Name,
			Uni:   user.UniID,
			Minat: minatID,
			Skill: skillID,
		}
		filter = append(filter, res)
	}

	// filter = removeUser(filter)

	return filter, nil
}

// func removeUser(users []model.UserFilter) []model.UserFilter {
// 	dUser := map[uuid.UUID]bool{}
// 	result := []model.UserFilter{}

// 	for _, user := range users {
// 		if dUser[user.ID] {
// 			continue
// 		}
// 		dUser[user.ID] = true
// 		result = append(result, user)
// 	}
// 	return result
// }

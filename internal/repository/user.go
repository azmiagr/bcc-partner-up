package repository

import (
	"intern-bcc/entity"
	"intern-bcc/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(u *entity.User) error
	GetUser(param model.UserParam) (entity.User, error)
	GetUserByName(name string) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(u *entity.User) error {
	return r.db.Omit("RoleID").Create(u).Error
}

func (u *UserRepository) GetUser(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where(&param).First(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepository) GetUserByName(name string) (*entity.User, error) {
	var user entity.User
	if err := u.db.Debug().Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// func (r *UserRepository) CreateUser(u entity.User) (entity.User,error) {
// 	err := r.db.Debug().Create(&u).Error
// 	if err != nil {
// 		return u, err
// 	}
// 	return u, nil
// }

//yang lama
// func (r *UserRepository) CreateUser(u *entity.User) error {
// 	return r.db.Omit("RoleID").Create(u).Error
//}

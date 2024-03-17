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
	UpdateUserPhoto(user entity.User, param model.UserParam) error
	UpdateProfile(id string, profileReq *model.UpdateProfile) (*entity.User, error)
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

func (u *UserRepository) UpdateUserPhoto(user entity.User, param model.UserParam) error {
	err := u.db.Debug().Where(&param).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) UpdateProfile(id string, profileReq *model.UpdateProfile) (*entity.User, error) {
	tx := u.db.Begin()
	var user entity.User
	err := u.db.Model(&user).Updates(entity.User{
		Name:       user.Name,
		UniID:      user.UniID,
		DistrictID: user.DistrictID,
		Minat:      user.Minat,
		Skill:      user.Skill,
	}).Error
	if err != nil {
		return nil, err
	}

	userParse := parseUpdateProfile(profileReq, &user)

	if err := tx.Debug().Model(&user).Save(&userParse).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &user, nil
}

func parseUpdateProfile(model *model.UpdateProfile, profile *entity.User) *entity.User {
	if model.Name != "" {
		profile.Name = model.Name
	}
	if model.District != 0 {
		profile.DistrictID = model.District
	}
	if model.Minat != nil {
		profile.Minat = model.Minat
	}
	if model.Skill != nil {
		profile.Skill = model.Skill
	}
	if model.Uni != 0 {
		profile.UniID = model.Uni
	}
	return profile
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

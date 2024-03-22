package repository

import (
	"fmt"
	"intern-bcc/entity"
	"intern-bcc/model"
	"intern-bcc/pkg/util"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	// CreateUser(u *entity.User) error
	CreateUser(u *entity.User) (*entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
	GetUserByName(name string) (*entity.User, error)
	UpdateUserPhoto(user entity.User, param model.UserParam) error
	UpdateProfile(id string, profileReq *model.UpdateProfile) (*entity.User, error)
	GetUsersByFilter(uniID uint, minatID []uint, skillID []uint) ([]entity.User, error)
	RecommendUser(userID uuid.UUID) ([]entity.Minat, error)
	FindRecommendUsers(userID uuid.UUID) ([]model.UserFilter, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(u *entity.User) (*entity.User, error) {
	if err := r.db.Debug().Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// func (r *UserRepository) CreateUser(u *entity.User) (error) {
// 	return r.db.Omit("RoleID").Create(u).Error
// }

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
	if err := u.db.Debug().Preload("Minat").Preload("Skill").Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) UpdateUserPhoto(user entity.User, param model.UserParam) error {
	// err := u.db.Debug().Where(&param).Updates(&user).Error
	err := u.db.Debug().Model(&entity.User{}).Where(param).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) UpdateProfile(id string, profileReq *model.UpdateProfile) (*entity.User, error) {
	tx := u.db.Begin()
	var user entity.User

	userID, err := uuid.Parse(id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	user.ID = userID
	err = tx.Model(&user).Select("name", "uni_id", "district_id").Updates(entity.User{
		Name:       profileReq.Name,
		UniID:      profileReq.Uni,
		DistrictID: profileReq.District,
	}).Error

	if err != nil {
		return nil, err
	}

	skill := make([]entity.Skill, 0, len(profileReq.Skill))
	minat := make([]entity.Minat, 0, len(profileReq.Minat))

	for _, v := range profileReq.Skill {
		skill = append(skill, entity.Skill{
			ID: v,
		})
	}

	for _, v := range profileReq.Minat {
		minat = append(minat, entity.Minat{
			ID: v,
		})
	}

	err = tx.Model(&user).Association("Skill").Replace(skill)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Model(&user).Association("Minat").Replace(minat)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &user, nil
}

func (u *UserRepository) GetUsersByFilter(uniID uint, minatID []uint, skillID []uint) ([]entity.User, error) {
	var users []entity.User
	query := u.db
	if uniID != 0 {
		query = query.Where("uni_id", uniID)
	}
	if len(minatID) > 0 {
		query = query.Preload("Minat").Joins("JOIN user_minat on users.id = user_minat.user_id").Where("user_minat.minat_id IN ?", minatID)
	}
	if len(skillID) > 0 {
		query = query.Preload("Skill").Joins("JOIN user_skill on users.id = user_skill.user_id").Where("user_skill.skill_id IN ?", skillID)
	}
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) RecommendUser(userID uuid.UUID) ([]entity.Minat, error) {
	var minat []entity.Minat
	if err := ur.db.Joins("JOIN user_minat ON minats.id = user_minat.minat_id").Where("user_minat.user_id IN (?)", userID).Find(&minat).Error; err != nil {
		return nil, err
	}
	return minat, nil
}

func (us *UserRepository) FindRecommendUsers(userID uuid.UUID) ([]model.UserFilter, error) {
	userMinat, err := us.RecommendUser(userID)
	if err != nil {
		return nil, err
	}

	combinations := Combinations(userMinat)

	var queries []string
	for _, combination := range combinations {
		if len(combination) == 0 {
			continue
		}
		query := "SELECT DISTINCT * FROM users JOIN user_minat ON users.id = user_minat.user_id WHERE "
		for i, minat := range combination {
			if i > 0 {
				if i < len(combination) {
					query += " AND "
				}
			}
			query += fmt.Sprintf("user_minat.minat_id=%v", minat.ID)
		}
		queries = append(queries, query)
	}

	var recommendUser []entity.User
	for _, query := range queries {
		var users []entity.User
		if err := us.db.Raw(query).Preload("Minat").Preload("Skill").Scan(&users).Error; err != nil {
			return nil, err
		}
		recommendUser = append(recommendUser, users...)
	}

	var res []model.UserFilter
	for _, recommend := range recommendUser {
		var minatID []uint
		var skillID []uint
		for _, minat := range recommend.Minat {
			minatID = append(minatID, minat.ID)
		}
		for _, skill := range recommend.Skill {
			skillID = append(skillID, skill.ID)
		}

		res = append(res, model.UserFilter{
			Name:  recommend.Name,
			Uni:   recommend.UniID,
			Minat: minatID,
			Skill: skillID,
		})
	}
	res = util.RemoveUser(res)
	return res, nil
}

func Combinations(minats []entity.Minat) [][]entity.Minat {
	var combination [][]entity.Minat
	n := len(minats)

	for i := 0; i < (1 << uint(n)); i++ {
		var combinations []entity.Minat
		for j := 0; j < n; j++ {
			if (i & (1 << uint(j))) > 0 {
				combinations = append(combinations, minats[j])
			}
		}
		combination = append(combination, combinations)
	}
	return combination
}

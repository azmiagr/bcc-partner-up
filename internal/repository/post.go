package repository

import (
	"intern-bcc/entity"
	"intern-bcc/model"

	"gorm.io/gorm"
)

type IPostRepository interface {
	CreatePost(post *entity.Post) (*entity.Post, error)
	UpdatePost(id string, postReq *model.UpdatePost) (*entity.Post, error)
}

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return &PostRepository{db}
}

func (pr *PostRepository) CreatePost(post *entity.Post) (*entity.Post, error) {
	if err := pr.db.Debug().Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (pr *PostRepository) UpdatePost(id string, postReq *model.UpdatePost) (*entity.Post, error) {
	up := pr.db.Begin()

	var post entity.Post
	if err := up.Debug().Where("id = ?", id).First(&post).Error; err != nil {
		up.Rollback()
		return nil, err
	}

	postParse := parseUpdate(postReq, &post)
	if err := up.Debug().Model(&post).Save(&postParse).Error; err != nil {
		up.Rollback()
		return nil, err
	}

	up.Commit()
	return &post, nil

}

func parseUpdate(model *model.UpdatePost, book *entity.Post) *entity.Post {
	if model.Title != "" {
		book.Title = model.Title
	}
	if model.Description != "" {
		book.Description = model.Description
	}
	return book
}

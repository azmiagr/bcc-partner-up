package service

import (
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
	"intern-bcc/model"

	"github.com/google/uuid"
)

type IPostService interface {
	CreatePost(postReq *model.CreatePost, id uuid.UUID) (*entity.Post, error)
	UpdatePost(id string, postReq *model.UpdatePost) (*entity.Post, error)
}

type PostService struct {
	PostRepository repository.IPostRepository
}

func NewPostService(postRepository repository.IPostRepository) IPostService {
	return &PostService{postRepository}
}

func (ps *PostService) CreatePost(postReq *model.CreatePost, id uuid.UUID) (*entity.Post, error) {
	post := &entity.Post{
		Title:       postReq.Title,
		Description: postReq.Description,
		UserID:      id,
	}

	post, err := ps.PostRepository.CreatePost(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (ps *PostService) UpdatePost(id string, postReq *model.UpdatePost) (*entity.Post, error) {
	post, err := ps.PostRepository.UpdatePost(id, postReq)
	if err != nil {
		return nil, err
	}

	return post, nil
}

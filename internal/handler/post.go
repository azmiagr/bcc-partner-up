package handler

import (
	"intern-bcc/entity"
	"intern-bcc/model"
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreatePost(ctx *gin.Context) {
	var postReq model.CreatePost
	user, _ := ctx.Get("user")
	userpost := user.(entity.User)

	if err := ctx.ShouldBindJSON(&postReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid request", err)
		return
	}

	post, err := r.service.Post.CreatePost(&postReq, userpost.ID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to create post", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "post created", post)
}

func (r *Rest) UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var postReq model.UpdatePost

	if err := ctx.ShouldBindJSON(&postReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Invalid request", err)
		return
	}

	book, err := r.service.Post.UpdatePost(id, &postReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update post", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Post updated", book)
}

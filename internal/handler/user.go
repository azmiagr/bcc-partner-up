package handler

import (
	// "intern-bcc/internal/service"

	"errors"
	"intern-bcc/entity"
	"intern-bcc/model"
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// func (r *Rest) Register(ctx *gin.Context) {
// 	param := model.UserRegister{}
// 	err := ctx.ShouldBindJSON(&param)
// 	if err != nil {
// 		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
// 		return
// 	}

// 	param, err = r.service.User.Register(&param)
// 	if err != nil {
// 		response.Error(ctx, http.StatusInternalServerError, "failed to register new user", err)
// 		return
// 	}

// 	response.Success(ctx, http.StatusCreated, "successfully register new user", nil)
// }

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	_, err = r.service.User.Register(&param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register new user", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "successfully register new user", nil)
}

func (r *Rest) Login(ctx *gin.Context) {
	param := model.UserLogin{}
	err := ctx.ShouldBindJSON(&param)

	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}
	result, err := r.service.User.Login(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to login user", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully login to system", result)

}

func (r *Rest) GetUserByName(ctx *gin.Context) {
	name := model.GetUserByNameRequest{}
	err := ctx.ShouldBindQuery(&name)

	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	user, err := r.service.User.GetUserByName(name.Name)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get user", err)
		return
	}

	responses := model.GetUserByNameResponse{
		Name:  user.Name,
		Uni:   user.UniID,
		Minat: user.Minat,
		Skill: user.Skill,
	}

	response.Success(ctx, http.StatusOK, "user found", responses)
}

func (r *Rest) UploadPhoto(ctx *gin.Context) {
	photo, err := ctx.FormFile("photo")
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to bind input", err)
		return
	}

	err = r.service.User.UploadPhoto(ctx, model.UploadPhoto{Photo: photo})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to upload data", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success upload photo", nil)
}

func (r *Rest) UpdateProfile(ctx *gin.Context) {
	id := ctx.Param("user_id")
	var profileReq model.UpdateProfile
	if err := ctx.ShouldBindJSON(&profileReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	profile, err := r.service.User.UpdateProfile(id, &profileReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update profile", err)
		return
	}

	res := model.UpdateProfileResponse{
		Name:     profile.Name,
		Uni:      profile.UniID,
		District: profile.DistrictID,
		Minat:    profileReq.Minat,
		Skill:    profileReq.Skill,
	}

	response.Success(ctx, http.StatusOK, "Profile Updated", res)
}

func (r *Rest) GetUsersByFilter(ctx *gin.Context) {
	var filter model.UserFilter

	if err := ctx.ShouldBindJSON(&filter); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	users, err := r.service.User.GetUsersByFilter(filter.Uni, filter.Minat, filter.Skill)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get user", err)
		return
	}

	res := make([]model.UserFilter, len(users))
	for i, user := range users {
		res[i] = model.UserFilter{
			Name:  user.Name,
			Uni:   user.Uni,
			Minat: user.Minat,
			Skill: user.Skill,
		}
	}

	response.Success(ctx, http.StatusOK, "user found", users)
}

func (r *Rest) GetRecommendUser(ctx *gin.Context) {
	userID := getUserID(ctx)
	recommendUser, err := r.service.User.RecommendUser(userID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get user", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully got user", recommendUser)
}

func getUserID(ctx *gin.Context) uuid.UUID {
	user, ok := ctx.Get("user")
	if !ok {
		response.Error(ctx, http.StatusUnauthorized, "User must login", errors.New("unauthorized"))
		return uuid.Nil
	}
	claims, ok := user.(entity.User)
	if !ok {
		response.Error(ctx, http.StatusUnauthorized, "User must login", errors.New("unauthorized"))
		return uuid.Nil
	}

	return claims.ID

}

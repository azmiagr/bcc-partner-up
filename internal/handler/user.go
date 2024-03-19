package handler

import (
	// "intern-bcc/internal/service"
	"intern-bcc/model"
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
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
	name := ctx.Param("name")

	user, err := r.service.User.GetUserByName(name)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get user", err)
		return
	}
	var profile model.UpdateProfileResponse
	responses := model.GetUserByNameResponse{
		Name:  user.Name,
		Uni:   user.UniID,
		Minat: profile.Minat,
		Skill: profile.Skill,
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
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid request", err)
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

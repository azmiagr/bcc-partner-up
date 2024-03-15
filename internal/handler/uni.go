package handler

import (
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetAllUni(ctx *gin.Context) {
	uni, err := r.service.Uni.GetAllUni()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get university", err)
		return
	}

	response.Success(ctx, http.StatusOK, "University found", uni)
}

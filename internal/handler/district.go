package handler

import (
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetAllDistrict(ctx *gin.Context) {
	uni, err := r.service.District.GetAllDistrict()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get district", err)
		return
	}

	response.Success(ctx, http.StatusOK, "District found", uni)
}

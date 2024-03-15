package handler

import (
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetAllMinat(ctx *gin.Context) {
	minat, err := r.service.Minat.GetAllMinat()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get minat", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Minat found", minat)
}

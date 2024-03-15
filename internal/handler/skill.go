package handler

import (
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetAllSkill(ctx *gin.Context) {
	skill, err := r.service.Skill.GetAllSkill()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get skill", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Skill found", skill)
}

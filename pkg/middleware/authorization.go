package middleware

import (
	"errors"
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *middleware) OnlyAdmin(ctx *gin.Context) {
	user, err := m.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		response.Error(ctx, http.StatusForbidden, "failed get login user", err)
		ctx.Abort()
		return
	}

	if user.RoleID != 1 {
		response.Error(ctx, http.StatusForbidden, "this endpoint cannot be access", errors.New("user don't have access"))
		ctx.Abort()
		return
	}

	ctx.Next()
}

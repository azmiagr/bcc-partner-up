package middleware

import (
	"errors"
	"intern-bcc/model"
	"intern-bcc/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.Error(ctx, http.StatusUnauthorized, "empty token", errors.New(""))
	}

	token := strings.Split(bearer, " ")[1]
	userid, err := m.jwtAuth.ValidateToken(token)

	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed to validate token", err)
		ctx.Abort()
		return
	}

	user, err := m.service.User.GetUser(model.UserParam{ID: userid})
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed to get user", err)
		ctx.Abort()
		return
	}

	ctx.Set("user", user)

	ctx.Next()
}

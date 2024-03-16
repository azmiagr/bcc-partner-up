package middleware

import (
	"intern-bcc/internal/service"
	"intern-bcc/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Interface interface {
	Timeout() gin.HandlerFunc
	Cors() gin.HandlerFunc
	AuthenticateUser(ctx *gin.Context)
	OnlyAdmin(ctx *gin.Context)
}

type middleware struct {
	jwtAuth jwt.Interface
	service *service.Service
}

func Init(jwtAuth jwt.Interface, service *service.Service) Interface {
	return &middleware{
		jwtAuth: jwtAuth,
		service: service,
	}
}

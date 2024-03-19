package handler

import (
	"errors"
	"fmt"
	"intern-bcc/internal/service"
	"intern-bcc/pkg/middleware"
	"intern-bcc/pkg/response"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	r.router.Use(r.middleware.Cors())
	r.router.POST("/register", r.Register)
	routerGroup := r.router.Group("api/v1")
	post := r.router.Group("user-post")
	r.router.Use(r.middleware.Timeout())

	routerGroup.GET("/login-user", r.middleware.AuthenticateUser, r.middleware.OnlyAdmin, testGetLoginUser)
	routerGroup.GET("/time-out", testTimeout)

	//r.middleware.OnlyAdmin,
	// routerGroup.GET("/uni", r.GetAllUni)
	// routerGroup.GET("/district", r.GetAllDistrict)
	// routerGroup.POST("/login", r.Login)
	// routerGroup.GET("/get-user/:name", r.GetUserByName)

	user := r.router.Group("/user")
	user.POST("/login", r.Login)
	user.GET("/skill", r.GetAllSkill)
	user.GET("/minat", r.GetAllMinat)
	user.GET("/uni", r.GetAllUni)
	user.GET("/district", r.GetAllDistrict)
	user.GET("/get-user/:name", r.middleware.AuthenticateUser, r.GetUserByName)
	user.POST("/profile/upload", r.middleware.AuthenticateUser, r.UploadPhoto)
	user.PATCH("/profile/update-profile/:user_id", r.middleware.AuthenticateUser, r.UpdateProfile)

	post.POST("/post", r.middleware.AuthenticateUser, r.CreatePost)
	post.PATCH("/update/:id/:user_id", r.UpdatePost)
}

func (r *Rest) Run() {
	// addr := os.Getenv("ADDRESS") // nanti ini command
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	r.router.Run(fmt.Sprintf(":%s", port))
}

func testTimeout(ctx *gin.Context) {
	time.Sleep(3 * time.Second)

	response.Success(ctx, http.StatusOK, "success", nil)
}

func testGetLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		response.Error(ctx, http.StatusUnauthorized, "test get login user", errors.New(" "))
		return
	}

	response.Success(ctx, http.StatusOK, "success", user)
}

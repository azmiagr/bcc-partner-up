package main

import (
	"intern-bcc/internal/handler"
	"intern-bcc/internal/repository"
	"intern-bcc/internal/service"
	"intern-bcc/pkg/bcrypt"
	"intern-bcc/pkg/config"
	"intern-bcc/pkg/database/mysql"
	"intern-bcc/pkg/jwt"
	"intern-bcc/pkg/middleware"
	"log"
)

func main() {
	config.LoadEnvironmet()

	db := mysql.ConnectDatabase()
	err := mysql.Migrate(db)

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)

	bcrypt := bcrypt.Init()

	jwt := jwt.Init()

	svc := service.NewService(repo, bcrypt, jwt)

	middleware := middleware.Init(jwt, svc)

	rest := handler.NewRest(svc, middleware)

	rest.MountEndpoint()

	rest.Run()

}

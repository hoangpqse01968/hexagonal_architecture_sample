package main

import (
	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/application/service"
	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/handler"
	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	v1Api := engine.Group("api/v1")

	userRepo := infrastructure.NewUserInfrastructure()
	userService := service.NewUserService(userRepo)
	getUserHandler := handler.NewGetUserHandler(userService)
	v1Api.GET("/user", getUserHandler.HandleGetUser)

	engine.Run(":8080")
}

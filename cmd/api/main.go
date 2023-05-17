package main

import (
	"context"

	"github.com/hoangpqse01968/hexagonal_architecture_sample/di"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	engine := gin.Default()
	v1Api := engine.Group("api/v1")

	getUserHandler := di.BuildUserHandler(ctx)
	v1Api.GET("/user", getUserHandler.HandleGetUser)

	engine.Run(":8080")
}

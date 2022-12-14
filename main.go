package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/pkg/handlers"
)

func main() {
	route := gin.Default()

	route.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "Hello"})
	})

	route.GET("/users/", handlers.UserHandler{}.Index)
	route.GET("/users/:user", handlers.UserHandler{}.Show)
	route.POST("/users/", handlers.UserHandler{}.Store)
	route.PUT("/users/:user", handlers.UserHandler{}.Update)
	route.PUT("/users/:user/password", handlers.UserHandler{}.UpdatePassword)
	route.DELETE("/users/:user", handlers.UserHandler{}.Destroy)

	err := route.Run(":8080")

	if err != nil {
		panic(err)
	}
}

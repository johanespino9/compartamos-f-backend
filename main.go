package main

import (
	"compartamos-backend/config"
	"compartamos-backend/controllers"
	"compartamos-backend/models"
	"compartamos-backend/repositories"
	"compartamos-backend/routes"
	"compartamos-backend/services"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        

        c.Next()
    }
}

func main() {
    config.ConnectDatabase()

    config.GetDB().AutoMigrate(&models.User{})

    userRepository := &repositories.UserRepository{DB: config.GetDB()}
    userService := services.UserService{UserRepository: *userRepository}
    userController := &controllers.UserController{UserService: userService}

    router := gin.Default()
    routes.UserRoutes(router, userController)

    router.Run(":8080")
}

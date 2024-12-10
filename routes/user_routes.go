package routes

import (
	"compartamos-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userController *controllers.UserController) {
    userGroup := router.Group("/users")
    {
        userGroup.GET("/", userController.GetUsers)
        userGroup.GET("/:id", userController.GetUser)
        userGroup.POST("/", userController.CreateUser)
        userGroup.PUT("/:id", userController.UpdateUser)
        userGroup.DELETE("/:id", userController.DeleteUser)
    }
}

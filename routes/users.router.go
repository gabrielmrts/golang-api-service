package routes

import (
	"github.com/gabrielmrts/golang-api-service/controllers"
	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.Engine) {

	var controller controllers.UserController

	r.GET("/users", controller.FindAll)

	r.POST("/users/auth", controller.Auth)

	r.POST("/users", controller.Add)

	r.GET("/users/:id", controller.FindOne)
}

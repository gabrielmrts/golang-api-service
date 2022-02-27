package routes

import (
	"github.com/gabrielmrts/golang-api-service/controllers"
	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.Engine) {

	var controller controllers.UserController

	r.GET("/users", func(c *gin.Context) {
		controller.FindAll(c)
	})

	r.POST("/users/auth", func(c *gin.Context) {
		controller.Auth(c)
	})

	r.POST("/users", func(c *gin.Context) {
		controller.Add(c)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		controller.FindOne(c)
	})
}

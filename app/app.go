package app

import (
	"github.com/gabrielmrts/golang-api-service/models"
	"github.com/gabrielmrts/golang-api-service/routes"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	models.Init()
	r := gin.Default()
	routes.MakeRoutes(r)
	r.Run()
}

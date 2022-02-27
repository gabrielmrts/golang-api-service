package main

import (
	"github.com/gabrielmrts/golang-api-service/models"
	"github.com/gabrielmrts/golang-api-service/routes"
	"github.com/gin-gonic/gin"
)

// Under Construction
//Just Ignore
func main() {
	models.Init()
	r := gin.Default()
	routes.MakeRoutes(r)
	r.Run()
}

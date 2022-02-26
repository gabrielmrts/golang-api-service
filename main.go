package main

import (
	"fmt"

	"github.com/gabrielmrts/golang-api-service/routes"
	"github.com/gabrielmrts/golang-api-service/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.MakeRoutes(r)
	r.Run()
	fmt.Print(service.Serv())
}

package routes

import (
	"github.com/gin-gonic/gin"
)

func MakeRoutes(r *gin.Engine) {
	userRoutes(r)
}

package routes

import (
	"api.dardasha.com/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/api/groups", controllers.CreateGroup())
}

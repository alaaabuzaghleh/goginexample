package controllers

import (
	"api.dardasha.com/services"
	"github.com/gin-gonic/gin"
)

func CreateGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.CreateGroup(c)
	}
}

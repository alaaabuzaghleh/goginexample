package main

import (
	"api.dardasha.com/db"
	"api.dardasha.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	db.ConnectToDB()

	//routes
	routes.UserRoute(router)

	router.Run("localhost:8000")
}

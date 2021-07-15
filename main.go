package main

import (
	"bean.com/web-service/database"
	"bean.com/web-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	route := gin.Default()
	routes.RouterAPI(route)

	// Listen and Server in http://localhost:8080
	route.Run()
}

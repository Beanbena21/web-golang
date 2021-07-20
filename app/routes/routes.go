package routes

import (
	"bean.com/web-service/app/controllers"
	"bean.com/web-service/middlewares"
	"github.com/gin-gonic/gin"
)

func RouterAPI(route *gin.Engine) {
	route.POST("/registerAccount", controllers.Register)
	route.POST("/loginAccount", controllers.Login)
	api := route.Group("/api")
	api.Use(middlewares.Authentication())
	{
		api.GET("/infoAccount/:id", controllers.InfoAcc)
	}
}

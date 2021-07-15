package routes

import (
	"bean.com/web-service/controllers"
	"github.com/gin-gonic/gin"
)

func RouterAPI(route *gin.Engine) {
	route.POST("/registerAccount", controllers.Register)
	route.POST("/loginAccount", controllers.Login)
	route.GET("/listAccount", controllers.ListAcc)
}

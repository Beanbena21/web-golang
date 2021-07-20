package controllers

import (
	"bean.com/web-service/database"
	"bean.com/web-service/middlewares"
	"bean.com/web-service/models"
	"bean.com/web-service/utils"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Context) {
	register := models.Register{}
	if err := g.ShouldBind(&register); err != nil {
		g.JSON(401, gin.H{
			"status":  401,
			"message": err.Error(),
		})
		return
	}
	value, err := database.HandleRegister(&register)
	if err != nil {
		g.AbortWithStatusJSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}
	g.JSON(200, gin.H{"status": 200, "message": "Register is success", "id": &value.InsertedID})
}

func Login(g *gin.Context) {
	login := models.Login{}
	if err := g.ShouldBind(&login); err != nil {
		g.JSON(401, gin.H{
			"status":  401,
			"message": err.Error(),
		})
		return
	}

	result, _ := database.HandleLogin(login.Email)
	if result.Email == "" {
		g.AbortWithStatusJSON(500, gin.H{
			"status":  500,
			"message": "Email not exist",
		})
		return
	}
	checkPass := utils.ComparePassword(result.User_password, login.User_password)
	if checkPass != nil {
		g.AbortWithStatusJSON(401, gin.H{
			"status":  401,
			"message": checkPass.Error(),
		})
		return
	}
	token, err := middlewares.GenerateToken(*result)
	if err != nil {
		g.AbortWithStatusJSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}
	g.JSON(200, gin.H{"status": 200, "message": "Log in success", "token": token})

}

func InfoAcc(g *gin.Context) {
	id := models.Id{}
	if err := g.ShouldBindUri(&id); err != nil {
		g.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	value, err := database.HandleInfo(id.ID)
	if err != nil {
		g.AbortWithStatusJSON(401, gin.H{
			"status":  401,
			"message": err.Error(),
		})
		return
	}
	g.JSON(200, gin.H{"status": 200, "info": value})
}

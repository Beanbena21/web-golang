package controllers

import (
	"fmt"

	"bean.com/web-service/database"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Context) {

}

func Login(g *gin.Context) {

}

func ListAcc(g *gin.Context) {
	fmt.Fprintln(g.Writer, "List account access")
	var users database.Users
	g.Bind(&users)
	g.JSON(200, gin.H{
		"user_name":     users.UserName,
		"user_password": users.UserPassword,
		"email":         users.Email,
	})
}

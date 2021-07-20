package middlewares

import (
	"strings"
	"time"

	"bean.com/web-service/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

func GenerateToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func Authentication() gin.HandlerFunc {
	return func(g *gin.Context) {
		requiredToken := g.GetHeader("Authorization")
		if len(requiredToken) == 0 {
			g.AbortWithStatusJSON(401, gin.H{
				"status":  401,
				"message": "Access Denied",
			})
			return
		}
		bearerToken := strings.Split(requiredToken, "Bearer ")
		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(bearerToken[1], claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				g.AbortWithStatusJSON(401, gin.H{
					"status":  401,
					"message": err.Error(),
				})
				return
			}
			g.AbortWithStatusJSON(500, gin.H{
				"status":  500,
				"message": err.Error(),
			})
			return
		}
		if !token.Valid {
			g.AbortWithStatusJSON(401, gin.H{
				"status":  401,
				"message": err.Error(),
			})
			return
		}
	}
}

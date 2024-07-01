package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenReal := ctx.GetHeader("Authorization")

		if tokenReal == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token not been found.",
			})
			return
		}

		secret := os.Getenv("JWT_PRIVATE_KEY")
		parseToken := strings.ReplaceAll(tokenReal, "Bearer. ", "")

		token, err := jwt.Parse(parseToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token is not valid sign",
			})
		}

		fmt.Println(token)

		ctx.Next()
	}
}

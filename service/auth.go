package service

import (
	"icecreambash/flika-backend/models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(user models.User) (string, error) {

	ttlToken, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * time.Duration(ttlToken)).Unix(),
		"iat": time.Now().Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_PRIVATE_KEY")))

	return token, err
}

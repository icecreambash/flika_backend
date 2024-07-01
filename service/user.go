package service

import (
	"fmt"
	"icecreambash/flika-backend/database"
	"icecreambash/flika-backend/models"
	"icecreambash/flika-backend/request"
	"icecreambash/flika-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(ctx *gin.Context, rq request.LoginUserRequest) {

	user := models.User{}

	database.Database.Where("email = ?", rq.Email).Find(&user)

	var isAuth bool = user != models.User{}

	if (user != models.User{}) {
		isAuth = utils.ComparePassword(user.Password, rq.Password)
	}

	if !isAuth {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Login or password is not correct",
		})
		return
	}

	token, err := GenerateJWTToken(user)

	fmt.Println(err)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

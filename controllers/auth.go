package controllers

import (
	"icecreambash/flika-backend/database"
	"icecreambash/flika-backend/models"
	"icecreambash/flika-backend/request"
	"icecreambash/flika-backend/service"
	"icecreambash/flika-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(ctx *gin.Context) {
	var loginRequest request.LoginUserRequest

	if err := ctx.ShouldBindWith(&loginRequest, binding.JSON); err != nil {
		utils.ExceptionError(ctx, err)
		return
	}

	service.LoginUser(ctx, loginRequest)

}

func RegisterUser(ctx *gin.Context) {
	var registerRequest request.RegisterUserRequest

	if err := ctx.ShouldBindWith(&registerRequest, binding.JSON); err != nil {
		utils.ExceptionError(ctx, err)
		return
	}

	var userModel = models.User{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
	}

	result := database.Database.Where("email = ?", userModel.Email).Find(&userModel)

	if result.RowsAffected > 0 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Email is reserved",
		})
		return
	}

	var password, _ = bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 12)

	userModel.Password = string(password)

	database.Database.Create(&userModel)

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully created user"})
}

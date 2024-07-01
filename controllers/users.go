package controllers

import (
	"icecreambash/flika-backend/database"
	"icecreambash/flika-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {

	var users []models.User

	database.Database.Find(&users)

	context.JSON(http.StatusGone, gin.H{"data": users})
}

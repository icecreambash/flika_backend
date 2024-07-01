package main

import (
	"icecreambash/flika-backend/controllers"
	"icecreambash/flika-backend/middleware"
	"icecreambash/flika-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Database *gorm.DB

func init() {
	utils.LoadEnv()
	utils.LoadDatabase()
}

func main() {
	r := gin.Default()

	authGroup := r.Group("auth")
	{
		authGroup.POST("login", controllers.LoginUser)
		authGroup.POST("register", controllers.RegisterUser)
	}

	users := r.Group("users").Use(middleware.AuthMiddleware())
	{
		users.GET("/", controllers.GetUsers)
	}

	r.Run("0.0.0.0:8881") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

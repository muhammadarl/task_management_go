package main

import (
	"net/http"
	"task_management_backend/config"
	"task_management_backend/controllers"
	"task_management_backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// database
	db := config.DatabaseConnection()
	db.AutoMigrate(&models.User{}, &models.Task{}, &models.TaskHistory{}, &models.UserHistory{}, &models.UserEmail{})
	config.CreateOwnerAccount(db)
	// controllers
	userController := controllers.UserControllers{DB: db}

	// router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello WOY!")
	})
	router.POST("/users/login", userController.Login)
	router.POST("/users/register", userController.Register)
	router.Static("/assets/img", "./assets/img")
	router.Run("10.200.2.171:80")
}

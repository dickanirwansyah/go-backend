package main

import (
	"bwaproject/handler"
	"bwaproject/users"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	databaseConnection := "root:root@tcp(127.0.0.1:3306)/bwaproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	usersRepository := users.NewRepository(db)
	usersService := users.NewService(usersRepository)
	userHandler := handler.NewUserHandler(usersService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	router.Run()
}

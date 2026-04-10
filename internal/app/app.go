package app

import (
	"github.com/MBFG9000/golang-jwt-auth/internal/config"
	"github.com/MBFG9000/golang-jwt-auth/internal/controller"
	"github.com/MBFG9000/golang-jwt-auth/internal/database"
	"github.com/MBFG9000/golang-jwt-auth/internal/repository"
	"github.com/MBFG9000/golang-jwt-auth/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
	godotenv.Load(".env")

	databaseConfig := config.DatabaseInit()

	dialect := database.GetConnection(databaseConfig)

	userRepo := repository.NewUserRepository(dialect)
	userService := service.NewUserService(userRepo)

	con := controller.NewController(userService)

	router := gin.Default()

	controller.SetupRoutes(router, con)

	router.Run()

}

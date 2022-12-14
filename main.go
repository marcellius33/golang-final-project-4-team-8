package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"os"
	"tokobelanja/controllers"
	"tokobelanja/database"
	"tokobelanja/database/seed"
	_ "tokobelanja/initializer"
	"tokobelanja/repositories"
	"tokobelanja/routers"
	"tokobelanja/services"
)

func init() {
	database.Connect()
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seeder":
			userRepository := repositories.NewUserRepository(database.GetDB())
			userSeed := seed.NewUserSeeder(userRepository)
			userSeed.Execute()
		}
	}
}

func main() {
	handleArgs()

	Routes := gin.Default()

	userRepository := repositories.NewUserRepository(database.GetDB())
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	routers.InitUserRoutes(Routes, userController)

	Routes.Run(os.Getenv("SERVER_PORT"))
}

package main

import (
	"flag"
	"os"
	"tokobelanja/controllers"
	"tokobelanja/database"
	"tokobelanja/database/seed"
	_ "tokobelanja/initializer"
	"tokobelanja/repositories"
	"tokobelanja/routers"
	"tokobelanja/services"

	"github.com/gin-gonic/gin"
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

	categoryRepository := repositories.NewCategoryRepository(database.GetDB())
	categoryService := services.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService)
	routers.InitCategoryRoutes(Routes, categoryController)

	productRepository := repositories.NewProductRepository(database.GetDB())
	productService := services.NewProductService(productRepository, categoryRepository)
	productController := controllers.NewProductController(productService)
	routers.InitProductRoutes(Routes, productController)

	transactionHistoryRepository := repositories.NewTransactionHistoryRepository(database.GetDB())
	transactionHistoryService := services.NewTransactionHistoryService(transactionHistoryRepository, productRepository, userRepository)
	transactionHistoryController := controllers.NewTransactionHistoryController(transactionHistoryService)
	routers.InitTransactionHistoryRoutes(Routes, transactionHistoryController)

	Routes.Run(os.Getenv("SERVER_PORT"))
}

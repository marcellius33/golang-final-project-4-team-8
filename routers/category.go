package routers

import (
	"github.com/gin-gonic/gin"
	"tokobelanja/controllers"
	"tokobelanja/middlewares"
)

func InitCategoryRoutes(Routes *gin.Engine, controller *controllers.CategoryController) {
	categoryRouter := Routes.Group("/categories")
	{
		categoryRouter.POST("", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), controller.CreateCategory)
		categoryRouter.GET("", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), controller.GetCategories)
		categoryRouter.PATCH("/:categoryId", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), controller.UpdateCategory)
		categoryRouter.DELETE("/:categoryId", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), controller.DeleteCategory)
	}
}

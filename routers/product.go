package routers

import (
	"tokobelanja/controllers"
	"tokobelanja/middlewares"

	"github.com/gin-gonic/gin"
)

func InitProductRoutes(Routes *gin.Engine, controller *controllers.ProductController) {
	productRouter := Routes.Group("/products")
	{
		productRouter.POST("", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), controller.CreateProduct)
		productRouter.GET("", middlewares.Authentication(), controller.GetProducts)
		productRouter.PUT("/:productId", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), controller.UpdateProduct)
		productRouter.DELETE("/:productId", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), controller.DeleteProduct)
	}
}

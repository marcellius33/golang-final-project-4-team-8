package routers

import (
	"github.com/gin-gonic/gin"
	"tokobelanja/controllers"
	"tokobelanja/middlewares"
)

func InitUserRoutes(Routes *gin.Engine, controller *controllers.UserController) {
	userRouter := Routes.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegisterController)
		userRouter.POST("/login", controller.UserLoginController)
		userRouter.PATCH("/topup", middlewares.Authentication(), controller.UserTopUpController)
	}
}

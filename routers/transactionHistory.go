package routers

import (
	"tokobelanja/controllers"
	"tokobelanja/middlewares"

	"github.com/gin-gonic/gin"
)

func InitTransactionHistoryRoutes(Routes *gin.Engine, controller *controllers.TransactionHistoryController) {
	transactionHistoryRouter := Routes.Group("/transactions")
	{
		transactionHistoryRouter.POST("", middlewares.Authentication(), middlewares.Authorization([]string{"customer"}), controller.CreateTransactionHistory)
		transactionHistoryRouter.GET("/my-transactions", middlewares.Authentication(), middlewares.Authorization([]string{"customer"}), controller.GetMyTransactionHistories)
		transactionHistoryRouter.GET("/user-transactions", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), controller.GetUserTransactionHistories)
	}
}

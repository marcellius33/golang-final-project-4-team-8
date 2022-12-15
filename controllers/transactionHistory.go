package controllers

import (
	"tokobelanja/helpers"
	"tokobelanja/params"
	"tokobelanja/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TransactionHistoryController struct {
	service services.TransactionHistoryService
}

func NewTransactionHistoryController(service services.TransactionHistoryService) *TransactionHistoryController {
	return &TransactionHistoryController{
		service: service,
	}
}

func (p *TransactionHistoryController) CreateTransactionHistory(c *gin.Context) {
	transactionHistoryRequest := params.CreateTransactionHistoryRequest{}
	if err := c.ShouldBindJSON(&transactionHistoryRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	userData, _ := c.MustGet("userData").(jwt.MapClaims)
	createTransactionHistory, err := p.service.CreateTransactionHistory(uint(userData["id"].(float64)), transactionHistoryRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessCreateTransactionHistoryResponse(createTransactionHistory, "Create TransactionHistory Success"))
}

func (p *TransactionHistoryController) GetMyTransactionHistories(c *gin.Context) {
	userData, _ := c.MustGet("userData").(jwt.MapClaims)
	transactionHistories, err := p.service.GetMyTransactionHistories(uint(userData["id"].(float64)))
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(transactionHistories, "Get My Transaction Histories Success"))
}

func (p *TransactionHistoryController) GetUserTransactionHistories(c *gin.Context) {
	transactionHistories, err := p.service.GetUserTransactionHistories()
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(transactionHistories, "Get User Transaction Histories Success"))
}

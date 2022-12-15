package services

import (
	"errors"
	"fmt"
	"tokobelanja/models"
	"tokobelanja/params"
	"tokobelanja/repositories"
)

type TransactionHistoryService interface {
	CreateTransactionHistory(userId uint, createTransactionHistoryRequest params.CreateTransactionHistoryRequest) (*params.CreateTransactionHistoryResponse, error)
	GetMyTransactionHistories(userId uint) (*[]params.GetMyTransactionHistoryResponse, error)
	GetUserTransactionHistories() (*[]params.GetUserTransactionHistoryResponse, error)
}

type transactionHistoryService struct {
	repository repositories.TransactionHistoryRepository
	productRepo repositories.ProductRepository
	userRepo repositories.UserRepository
}

func NewTransactionHistoryService(
	repository repositories.TransactionHistoryRepository,
	productRepo repositories.ProductRepository,
	userRepo repositories.UserRepository,
	) TransactionHistoryService {
	return &transactionHistoryService{
		repository: repository,
		productRepo: productRepo,
		userRepo: userRepo,
	}
}


func (c transactionHistoryService) CreateTransactionHistory(userId uint, createTransactionHistoryRequest params.CreateTransactionHistoryRequest) (*params.CreateTransactionHistoryResponse, error) {
	product, err := c.productRepo.FindProductByID(createTransactionHistoryRequest.ProductID)
	if err != nil {
		return &params.CreateTransactionHistoryResponse{}, err
	}

	if createTransactionHistoryRequest.Quantity > product.Stock{
		return &params.CreateTransactionHistoryResponse{}, errors.New("Product stock is not enough")
	}

	user, err := c.userRepo.FindUserByID(userId)
	if err != nil {
		return &params.CreateTransactionHistoryResponse{}, err
	}

	if product.Price*createTransactionHistoryRequest.Quantity > user.Balance {
		return &params.CreateTransactionHistoryResponse{}, errors.New("Your balance is insufficient")
	}

	newTransactionHistory := models.TransactionHistory{
		ProductID: createTransactionHistoryRequest.ProductID,
		UserID: userId,
		Quantity: createTransactionHistoryRequest.Quantity,
		TotalPrice: createTransactionHistoryRequest.Quantity*product.Price,
	}
	_, err = c.repository.CreateTransactionHistory(&newTransactionHistory)
	if err != nil {
		return &params.CreateTransactionHistoryResponse{}, err
	}

	product.Stock = product.Stock-newTransactionHistory.Quantity
	fmt.Println(product.Stock)
	fmt.Println(product)
	_, err = c.productRepo.UpdateProduct(product.ID, product)
	if err != nil {
		return &params.CreateTransactionHistoryResponse{}, err
	}

	user.Balance = user.Balance-newTransactionHistory.TotalPrice
	_, err = c.userRepo.UpdateUser(userId, user)
	if err != nil {
		return &params.CreateTransactionHistoryResponse{}, err
	}

	resp := params.ParseToCreateTransactionHistoryResponse(&newTransactionHistory, product)

	return &resp, err
}

func (c transactionHistoryService) GetMyTransactionHistories(userId uint) (*[]params.GetMyTransactionHistoryResponse, error) {
	var transactionHistories []models.TransactionHistory
	_, err := c.repository.GetMyTransactionHistories(userId, &transactionHistories)
	if err != nil {
		return &[]params.GetMyTransactionHistoryResponse{}, err
	}
	resp := params.ParseToGetMyTransactionResponse(&transactionHistories)

	return &resp, err
}

func (c transactionHistoryService) GetUserTransactionHistories() (*[]params.GetUserTransactionHistoryResponse, error) {
	var transactionHistories []models.TransactionHistory
	_, err := c.repository.GetUserTransactionHistories(&transactionHistories)
	if err != nil {
		return &[]params.GetUserTransactionHistoryResponse{}, err
	}
	resp := params.ParseToGetUserTransactionResponse(&transactionHistories)

	return &resp, err
}
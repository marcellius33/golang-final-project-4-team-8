package params

import (
	"time"
	"tokobelanja/models"
)

type CreateTransactionHistoryResponse struct {
	TotalPrice uint      `json:"total_price"`
	Quantity   uint      `json:"quantity"`
	Title      string 	 `json:"product_title"`
	
}

type CreateTransactionHistoryRequest struct {
	ProductID  uint      `json:"product_id" binding:"required"`
	Quantity   uint      `json:"quantity" binding:"required"`
}

type GetMyTransactionHistoryResponse struct {
	ID         uint      `json:"id"`
	ProductID  uint      `json:"product_id"`
	UserID     uint      `json:"user_id"`
	Quantity   uint      `json:"quantity"`
	TotalPrice uint      `json:"total_price"`
	Product    GetTransactionProductResponse `json:"Product"`
}

type GetUserTransactionHistoryResponse struct {
	ID         uint      `json:"id"`
	ProductID  uint      `json:"product_id"`
	UserID     uint      `json:"user_id"`
	Quantity   uint      `json:"quantity"`
	TotalPrice uint      `json:"total_price"`
	Product    GetTransactionProductResponse `json:"Product"`
	User	   GetTransactionUserResponse 	 `json:"User"`
}

type GetTransactionProductResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      uint      `json:"stock"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetTransactionUserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Balance   uint      `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func ParseToCreateTransactionHistoryResponse(transaction *models.TransactionHistory, product *models.Product) CreateTransactionHistoryResponse {
	return CreateTransactionHistoryResponse{
		TotalPrice:	transaction.TotalPrice,
		Quantity:	transaction.Quantity,
		Title: 		product.Title,
	}
}

func ParseToGetMyTransactionResponse(transactions *[]models.TransactionHistory) []GetMyTransactionHistoryResponse {
	var responses []GetMyTransactionHistoryResponse
	for _, transaction := range *transactions {
		responses = append(responses, GetMyTransactionHistoryResponse{
			ID:			transaction.ID,
			ProductID:	transaction.ProductID,
			UserID:		transaction.UserID,
			Quantity:	transaction.Quantity,
			TotalPrice:	transaction.TotalPrice,
			Product:   ParseToGetTransactionProductResponse(&transaction.Product),
		})
	}

	return responses
}

func ParseToGetUserTransactionResponse(transactions *[]models.TransactionHistory) []GetUserTransactionHistoryResponse {
	var responses []GetUserTransactionHistoryResponse
	for _, transaction := range *transactions {
		responses = append(responses, GetUserTransactionHistoryResponse{
			ID:			transaction.ID,
			ProductID:	transaction.ProductID,
			UserID:		transaction.UserID,
			Quantity:	transaction.Quantity,
			TotalPrice:	transaction.TotalPrice,
			Product:   ParseToGetTransactionProductResponse(&transaction.Product),
			User:   ParseToGetTransactionUserResponse(&transaction.User),
		})
	}

	return responses
}

func ParseToGetTransactionProductResponse(product *models.Product) GetTransactionProductResponse {
	return GetTransactionProductResponse{
		ID:			product.ID,
		Title:		product.Title,
		Price:		product.Price,
		Stock:		product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:	product.CreatedAt,
		UpdatedAt:	product.UpdatedAt,
	}
}

func ParseToGetTransactionUserResponse(user *models.User) GetTransactionUserResponse {
	return GetTransactionUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}
}

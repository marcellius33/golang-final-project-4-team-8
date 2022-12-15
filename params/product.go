package params

import (
	"time"
	"tokobelanja/models"
)

type CreateProductResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Price      uint   `json:"price"`
	Stock      uint   `json:"stock"`
	CategoryID uint   `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateProductRequest struct {
	Title      string `json:"title" binding:"required"`
	Price      uint   `json:"price" binding:"required"`
	Stock      uint   `json:"stock" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

type GetProductResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Price      uint   `json:"price"`
	Stock      uint   `json:"stock"`
	CategoryID uint   `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type UpdateProductResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Price      uint   `json:"price"`
	Stock      uint   `json:"stock"`
	CategoryID uint   `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UpdateProductRequest struct {
	Title      string `json:"title" binding:"required"`
	Price      uint   `json:"price" binding:"required"`
	Stock      uint   `json:"stock" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

func ParseToCreateProductResponse(product *models.Product) CreateProductResponse {
	return CreateProductResponse{
		ID:			product.ID,
		Title:		product.Title,
		Price:		product.Price,
		Stock:		product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:	product.CreatedAt,
	}
}

func ParseToGetProductResponse(products *[]models.Product) []GetProductResponse {
	var responses []GetProductResponse
	for _, product := range *products {
		responses = append(responses, GetProductResponse{
			ID:			product.ID,
			Title:		product.Title,
			Price:		product.Price,
			Stock:		product.Stock,
			CategoryID: product.CategoryID,
			CreatedAt:	product.CreatedAt,
		})
	}

	return responses
}

func ParseToUpdateProductResponse(product *models.Product) UpdateProductResponse {
	return UpdateProductResponse{
		ID:			product.ID,
		Title:		product.Title,
		Price:		product.Price,
		Stock:		product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:	product.CreatedAt,
		UpdatedAt:	product.UpdatedAt,
	}
}

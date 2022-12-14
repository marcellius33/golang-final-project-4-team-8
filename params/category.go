package params

import (
	"time"
	"tokobelanja/models"
)

type CreateCategoryResponse struct {
	ID                uint      `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount uint      `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

type CreateCategoryRequest struct {
	Type string `json:"type" binding:"required"`
}

type GetCategoryResponse struct {
	ID                uint                         `json:"id"`
	Type              string                       `json:"type"`
	SoldProductAmount uint                         `json:"sold_product_amount"`
	CreatedAt         time.Time                    `json:"created_at"`
	UpdatedAt         time.Time                    `json:"updated_at"`
	Products          []GetCategoryProductResponse `json:"Products"`
}

type GetCategoryProductResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Price     uint      `json:"price"`
	Stock     uint      `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateCategoryResponse struct {
	ID                uint      `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount uint      `json:"sold_product_amount"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type UpdateCategoryRequest struct {
	Type string `json:"type" binding:"required"`
}

func ParseToCreateCategoryResponse(category *models.Category) CreateCategoryResponse {
	return CreateCategoryResponse{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		CreatedAt:         category.CreatedAt,
	}
}

func ParseToGetCategoryResponse(categories *[]models.Category) []GetCategoryResponse {
	var responses []GetCategoryResponse
	for _, category := range *categories {
		responses = append(responses, GetCategoryResponse{
			ID:                category.ID,
			Type:              category.Type,
			SoldProductAmount: category.SoldProductAmount,
			CreatedAt:         category.CreatedAt,
			UpdatedAt:         category.UpdatedAt,
			Products:          ParseToGetCategoryProductResponse(&category.Products),
		})
	}

	return responses
}

func ParseToGetCategoryProductResponse(products *[]models.Product) []GetCategoryProductResponse {
	var responses []GetCategoryProductResponse
	for _, product := range *products {
		responses = append(responses, GetCategoryProductResponse{
			ID:        product.ID,
			Title:     product.Title,
			Price:     product.Price,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		})
	}
	return responses
}

func ParseToUpdateCategoryResponse(category *models.Category) UpdateCategoryResponse {
	return UpdateCategoryResponse{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		UpdatedAt:         category.UpdatedAt,
	}
}

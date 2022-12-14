package services

import (
	"time"
	"tokobelanja/models"
	"tokobelanja/params"
	"tokobelanja/repositories"
)

type CategoryService interface {
	CreateCategory(createCategoryRequest params.CreateCategoryRequest) (*params.CreateCategoryResponse, error)
	GetCategories() (*[]params.GetCategoryResponse, error)
	UpdateCategory(id uint, updateCategoryRequest params.UpdateCategoryRequest) (*params.UpdateCategoryResponse, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	repository repositories.CategoryRepository
}

func (c categoryService) CreateCategory(createCategoryRequest params.CreateCategoryRequest) (*params.CreateCategoryResponse, error) {
	newCategory := models.Category{
		Type:              createCategoryRequest.Type,
		SoldProductAmount: 0,
	}

	_, err := c.repository.CreateCategory(&newCategory)
	if err != nil {
		return &params.CreateCategoryResponse{}, err
	}
	resp := params.ParseToCreateCategoryResponse(&newCategory)

	return &resp, err
}

func (c categoryService) GetCategories() (*[]params.GetCategoryResponse, error) {
	var categories []models.Category
	_, err := c.repository.GetCategories(&categories)
	if err != nil {
		return &[]params.GetCategoryResponse{}, err
	}
	resp := params.ParseToGetCategoryResponse(&categories)

	return &resp, err
}

func (c categoryService) UpdateCategory(id uint, updateCategoryRequest params.UpdateCategoryRequest) (*params.UpdateCategoryResponse, error) {
	categoryModel, err := c.repository.FindCategoryByID(id)
	if err != nil {
		return &params.UpdateCategoryResponse{}, err
	}

	categoryModel.Type = updateCategoryRequest.Type
	categoryModel.UpdatedAt = time.Now()
	category, err := c.repository.UpdateCategory(id, categoryModel)
	if err != nil {
		return &params.UpdateCategoryResponse{}, err
	}

	resp := params.ParseToUpdateCategoryResponse(category)
	return &resp, err
}

func (c categoryService) DeleteCategory(id uint) error {
	return c.repository.DeleteCategory(id)
}

func NewCategoryService(repository repositories.CategoryRepository) CategoryService {
	return &categoryService{
		repository: repository,
	}
}

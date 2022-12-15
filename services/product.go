package services

import (
	"time"
	"tokobelanja/models"
	"tokobelanja/params"
	"tokobelanja/repositories"
)

type ProductService interface {
	CreateProduct(createProductRequest params.CreateProductRequest) (*params.CreateProductResponse, error)
	GetProducts() (*[]params.GetProductResponse, error)
	UpdateProduct(id uint, updateProductRequest params.UpdateProductRequest) (*params.UpdateProductResponse, error)
	DeleteProduct(id uint) error
}

type productService struct {
	repository repositories.ProductRepository
	categoryRepo repositories.CategoryRepository
}

func NewProductService(repository repositories.ProductRepository, categoryRepo repositories.CategoryRepository) ProductService {
	return &productService{
		repository: repository,
		categoryRepo: categoryRepo,
	}
}


func (c productService) CreateProduct(createProductRequest params.CreateProductRequest) (*params.CreateProductResponse, error) {
	_, err := c.categoryRepo.FindCategoryByID(createProductRequest.CategoryID)
	if err != nil {
		return &params.CreateProductResponse{}, err
	}
	
	newProduct := models.Product{
		Title: createProductRequest.Title,
		Price: createProductRequest.Price,
		Stock: createProductRequest.Stock,
		CategoryID: createProductRequest.CategoryID,
	}

	_, err = c.repository.CreateProduct(&newProduct)
	if err != nil {
		return &params.CreateProductResponse{}, err
	}
	resp := params.ParseToCreateProductResponse(&newProduct)

	return &resp, err
}

func (c productService) GetProducts() (*[]params.GetProductResponse, error) {
	var products []models.Product
	_, err := c.repository.GetProducts(&products)
	if err != nil {
		return &[]params.GetProductResponse{}, err
	}
	resp := params.ParseToGetProductResponse(&products)

	return &resp, err
}

func (c productService) UpdateProduct(id uint, updateProductRequest params.UpdateProductRequest) (*params.UpdateProductResponse, error) {
	ProductModel, err := c.repository.FindProductByID(id)
	if err != nil {
		return &params.UpdateProductResponse{}, err
	}

	_, err = c.categoryRepo.FindCategoryByID(updateProductRequest.CategoryID)
	if err != nil {
		return &params.UpdateProductResponse{}, err
	}

	ProductModel.Title = updateProductRequest.Title
	ProductModel.Price = updateProductRequest.Price
	ProductModel.Stock = updateProductRequest.Stock
	ProductModel.CategoryID = updateProductRequest.CategoryID
	ProductModel.UpdatedAt = time.Now()
	Product, err := c.repository.UpdateProduct(id, ProductModel)
	if err != nil {
		return &params.UpdateProductResponse{}, err
	}

	resp := params.ParseToUpdateProductResponse(Product)
	return &resp, err
}

func (c productService) DeleteProduct(id uint) error {
	return c.repository.DeleteProduct(id)
}
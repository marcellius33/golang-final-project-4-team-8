package controllers

import (
	"strconv"
	"tokobelanja/helpers"
	"tokobelanja/params"
	"tokobelanja/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	productRequest := params.CreateProductRequest{}
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	createProduct, err := p.service.CreateProduct(productRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessCreateResponse(createProduct, "Create Product Success"))
}

func (p *ProductController) GetProducts(c *gin.Context) {
	products, err := p.service.GetProducts()
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(products, "Get Products Success"))
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	productRequest := params.UpdateProductRequest{}
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}
	productId, err := strconv.Atoi(c.Param("productId"))

	updateProduct, err := p.service.UpdateProduct(uint(productId), productRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(updateProduct, "Update Product Success"))
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}
	err = p.service.DeleteProduct(uint(productId))
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.DeleteSuccess("Product has been successfully deleted"))
}
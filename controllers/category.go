package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tokobelanja/helpers"
	"tokobelanja/params"
	"tokobelanja/services"
)

type CategoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	categoryRequest := params.CreateCategoryRequest{}
	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	createCategory, err := cc.service.CreateCategory(categoryRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessCreateResponse(createCategory, "Create Category Success"))
}

func (cc *CategoryController) GetCategories(c *gin.Context) {
	categories, err := cc.service.GetCategories()
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(categories, "Get Categories Success"))
}

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	categoryRequest := params.UpdateCategoryRequest{}
	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}
	categoryId, err := strconv.Atoi(c.Param("categoryId"))

	updatedCategory, err := cc.service.UpdateCategory(uint(categoryId), categoryRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(updatedCategory, "Update Category Success"))
}

func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}
	err = cc.service.DeleteCategory(uint(categoryId))
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.DeleteSuccess("Category has been successfully deleted"))
}

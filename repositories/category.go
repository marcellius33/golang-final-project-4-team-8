package repositories

import (
	"gorm.io/gorm"
	"tokobelanja/models"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) (*models.Category, error)
	GetCategories(categories *[]models.Category) (*[]models.Category, error)
	UpdateCategory(id uint, category *models.Category) (*models.Category, error)
	DeleteCategory(id uint) error
	FindCategoryByID(id uint) (*models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func (c categoryRepository) CreateCategory(category *models.Category) (*models.Category, error) {
	return category, c.db.Create(category).Error
}

func (c categoryRepository) GetCategories(categories *[]models.Category) (*[]models.Category, error) {
	err := c.db.Preload("Products").Find(&categories).Error
	return categories, err
}

func (c categoryRepository) UpdateCategory(id uint, updateCategory *models.Category) (*models.Category, error) {
	category := updateCategory
	err := c.db.Model(&category).Where("id = ?", id).Updates(updateCategory).Error
	return category, err
}

func (c categoryRepository) DeleteCategory(id uint) error {
	return c.db.Where("id = ?", id).Delete(&models.Category{}).Error
}

func (c categoryRepository) FindCategoryByID(id uint) (*models.Category, error) {
	category := models.Category{}
	err := c.db.Where("id = ?", id).First(&category).Error
	return &category, err
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

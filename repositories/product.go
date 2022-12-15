package repositories

import (
	"tokobelanja/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProducts(products *[]models.Product) (*[]models.Product, error)
	UpdateProduct(id uint, product *models.Product) (*models.Product, error)
	DeleteProduct(id uint) error
	FindProductByID(id uint) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (c productRepository) CreateProduct(Product *models.Product) (*models.Product, error) {
	return Product, c.db.Create(Product).Error
}

func (c productRepository) GetProducts(products *[]models.Product) (*[]models.Product, error) {
	err := c.db.Find(&products).Error
	return products, err
}

func (c productRepository) UpdateProduct(id uint, updateProduct *models.Product) (*models.Product, error) {
	product := updateProduct
	err := c.db.Model(&product).Where("id = ?", id).Updates(updateProduct).Error
	return product, err
}

func (c productRepository) DeleteProduct(id uint) error {
	return c.db.Where("id = ?", id).Delete(&models.Product{}).Error
}

func (c productRepository) FindProductByID(id uint) (*models.Product, error) {
	product := models.Product{}
	err := c.db.Where("id = ?", id).First(&product).Error
	return &product, err
}

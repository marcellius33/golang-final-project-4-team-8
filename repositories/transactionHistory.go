package repositories

import (
	"tokobelanja/models"

	"gorm.io/gorm"
)

type TransactionHistoryRepository interface {
	CreateTransactionHistory(transactionHistory *models.TransactionHistory) (*models.TransactionHistory, error)
	GetMyTransactionHistories(userId uint, transactionHistories *[]models.TransactionHistory) (*[]models.TransactionHistory, error)
	GetUserTransactionHistories(transactionHistories *[]models.TransactionHistory) (*[]models.TransactionHistory, error)
}

type transactionHistoryRepository struct {
	db *gorm.DB
}

func NewTransactionHistoryRepository(db *gorm.DB) TransactionHistoryRepository {
	return &transactionHistoryRepository{
		db: db,
	}
}

func (c transactionHistoryRepository) CreateTransactionHistory(transactionHistory *models.TransactionHistory) (*models.TransactionHistory, error) {
	return transactionHistory, c.db.Create(transactionHistory).Error
}

func (c transactionHistoryRepository) GetMyTransactionHistories(userId uint, transactionHistories *[]models.TransactionHistory) (*[]models.TransactionHistory, error) {
	err := c.db.Preload("Product").Find(&transactionHistories).Error
	return transactionHistories, err
}

func (c transactionHistoryRepository) GetUserTransactionHistories(transactionHistories *[]models.TransactionHistory) (*[]models.TransactionHistory, error) {
	err := c.db.Preload("Product").Preload("User").Find(&transactionHistories).Error
	return transactionHistories, err
}
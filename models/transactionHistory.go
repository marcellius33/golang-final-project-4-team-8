package models

import "time"

// Quantity: not null
// TotalPrice: not null

type TransactionHistory struct {
	ID         uint      `json:"id" gorm:"primaryKey;type:integer"`
	ProductID  uint      `json:"product_id"`
	UserID     uint      `json:"user_id"`
	Quantity   uint      `json:"quantity" gorm:"type:integer;not null"`
	TotalPrice uint      `json:"total_price" gorm:"type:integer; not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Product    Product
	User       User
}

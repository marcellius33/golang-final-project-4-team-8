package models

import "time"

// Type: not null

type Category struct {
	ID                uint      `json:"id" gorm:"primaryKey;type:integer"`
	Type              string    `json:"type" gorm:"type:varchar(255);not null"`
	SoldProductAmount uint      `json:"sold_product_amount" gorm:"type:integer"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Products          []Product `json:"Products"`
}

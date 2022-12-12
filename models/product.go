package models

import "time"

// Title: not null
// Stock: not null, > 5
// Price: not null, > 0, < 50000000

type Product struct {
	ID         uint   `json:"id" gorm:"primaryKey;type:integer"`
	Title      string `json:"title" gorm:"type:varchar(255);not null"`
	Price      uint   `json:"price" gorm:"type:integer;not null;check: price >= 0 and price <= 50000000"`
	Stock      uint   `json:"stock" gorm:"type:integer;not null;check: stock >= 5"`
	CategoryID uint   `json:"category_id"`
	Category   Category
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

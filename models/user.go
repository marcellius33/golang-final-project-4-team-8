package models

import "time"

// Email : format valid, unique index, not null
// FullName: not null
// Balance: not null, max 100.000.000, min 0
// Password: not null, min 6
// Role: not null, only admin or customer

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;type:integer"`
	FullName  string    `json:"full_name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Role      string    `json:"role" gorm:"type:varchar(255); not null"`
	Balance   uint      `json:"balance" gorm:"type:integer;check:balance >= 0 and balance <= 100000000"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

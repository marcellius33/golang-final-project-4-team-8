package params

import (
	"time"
	"tokobelanja/models"
)

type UserRegisterResponse struct {
	ID        uint      `json:"id"  example:"1"`
	FullName  string    `json:"full_name" example:"curry"`
	Email     string    `json:"email" example:"curry@gmail.com"`
	Password  string    `json:"password" example:"curry123"`
	Balance   uint      `json:"balance" example:"0"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserTopUpRequest struct {
	Balance uint `json:"balance" example:"0"`
}

func ParseToCreateUserResponse(user *models.User) UserRegisterResponse {
	return UserRegisterResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}
}

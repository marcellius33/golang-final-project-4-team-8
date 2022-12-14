package seed

import (
	"golang.org/x/crypto/bcrypt"
	"tokobelanja/models"
	"tokobelanja/repositories"
)

type UserSeeder struct {
	repository repositories.UserRepository
}

func NewUserSeeder(repository repositories.UserRepository) *UserSeeder {
	return &UserSeeder{
		repository: repository,
	}
}

func (u *UserSeeder) Execute() {
	pwHash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	admin := models.User{
		FullName: "admin",
		Email:    "admin@gmail.com",
		Password: string(pwHash),
		Role:     "admin",
		Balance:  0,
	}

	_, _ = u.repository.CreateUser(&admin)
}

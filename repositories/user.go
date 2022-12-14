package repositories

import (
	"gorm.io/gorm"
	"tokobelanja/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, user *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) CreateUser(user *models.User) (*models.User, error) {
	return user, u.db.Create(user).Error
}

func (u userRepository) FindUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u userRepository) FindUserByID(id uint) (*models.User, error) {
	user := models.User{}
	err := u.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u userRepository) UpdateUser(id uint, updatedUser *models.User) (*models.User, error) {
	user := updatedUser
	err := u.db.Model(&user).Where("id = ?", id).Updates(updatedUser).Error
	return user, err
}

func (u userRepository) DeleteUser(id uint) error {
	user := models.User{}
	return u.db.Where("id = ?", id).Delete(&user).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

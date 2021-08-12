package repository

import (
	"bookstore/models"

	"gorm.io/gorm"
)

// UserRepository --> CRUD of user
type UserRepository interface {
	AddUser(models.User) (models.User, error)
	GetUser(int) (models.User, error)
	GetUserByEmail(string) (models.User, error)
	GetAllUser() ([]models.User, error)
	GetUserProducts(int) ([]models.Order, error)
	UpdateUser(models.User) (models.User, error)
	DeleteUser(models.User) (models.User, error)
}

type userRepository struct {
	connection *gorm.DB
}

// NewUserRepository --> returs new user repository
func NewUserRepository() UserRepository {
	return &userRepository{
		connection: DB(),
	}
}

func (db *userRepository) AddUser(user models.User) (models.User, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) GetUser(userID int) (user models.User, err error) {
	return user, db.connection.First(&user, userID).Error
}

func (db *userRepository) GetUserByEmail(email string) (user models.User, err error) {
	return user, db.connection.First(&user, "email=?", email).Error
}

func (db *userRepository) GetAllUser() (users []models.User, err error) {
	return users, db.connection.Find(&users).Error
}

func (db *userRepository) UpdateUser(user models.User) (models.User, error) {
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}

	return user, db.connection.Model(&user).Updates(&user).Error
}

func (db *userRepository) DeleteUser(user models.User) (models.User, error) {
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}

	return user, db.connection.Delete(&user).Error
}

func (db *userRepository) GetUserProducts(userID int) (orders []models.Order, err error) {
	return orders, db.connection.Where("user_id=?", userID).Set("gorm:auto_preload", true).Find(&orders).Error
}

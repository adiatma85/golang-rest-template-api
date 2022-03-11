package repository

import (
	"strconv"

	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/models"
	"github.com/adiatma85/go-tutorial-gorm/pkg/crypto"
)

// Local variable
var (
	err            error
	userRepository *UserRepository
)

// Contract of User Repository
type UserRepositoryInterface interface {
	CreateUser(user models.User) (models.User, error)
}

// Struct to implements contract or interface
type UserRepository struct{}

// Func to return User Repository instance
func GetUserRepository() UserRepositoryInterface {
	if userRepository == nil {
		userRepository = &UserRepository{}
	}
	return userRepository
}

// Func to Create User
func (repo *UserRepository) CreateUser(user models.User) (models.User, error) {
	user.Password, err = crypto.HashAndSalt([]byte(user.Password))
	if err != nil {
		return models.User{}, err
	}
	Create(&user)
	return user, nil
}

// Func to get All User
// func GetAllUser() (*[]models.User, error) {
// 	// var users []models.User
// 	// err := Find
// }

// Func to Get User By Id
func (repo *UserRepository) GetUserById(userId string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.ID, _ = strconv.ParseUint(userId, 10, 64)
	_, err := First(&where, &user, []string{})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

package repository

import (
	"fmt"
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
	Create(user models.User) (models.User, error)
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
func (repo *UserRepository) Create(user models.User) (models.User, error) {
	user.Password, err = crypto.HashAndSalt([]byte(user.Password))
	if err != nil {
		return models.User{}, err
	}
	Create(&user)
	return user, nil
}

// Func to get All User without Pagination
func (repo *UserRepository) GetAll() (*[]models.User, error) {
	var users []models.User
	err := Find(&models.User{}, &users, []string{""}, "id asc")
	return &users, err
}

// Func to get Query of WHERE withoud Pagination
func (repo *UserRepository) Query(q *models.User) (*[]models.User, error) {
	var users []models.User
	err := Find(&q, &users, []string{""}, "id asc")
	return &users, err
}

// Func to Get User By Id
func (repo *UserRepository) GetById(userId string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.ID, _ = strconv.ParseUint(userId, 10, 64)
	_, err := First(&where, &user, []string{})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Function to update user according to user schema defined
func (repo *UserRepository) Update(user *models.User) error {
	if user.Password != "" {
		user.Password, err = crypto.HashAndSalt([]byte(user.Password))
		if err != nil {
			return err
		}
	} else {
		var tempUser *models.User
		tempUser, err = repo.GetById(fmt.Sprint(user.ID))
		user.Password = tempUser.Password
		if err != nil {
			return err
		}
	}
	return nil
}

// Delete User By Model defined in controller
func (repo *UserRepository) Delete(user *models.User) error {
	_, err = DeleteByModel(user)
	if err != nil {
		return err
	}
	return nil
}

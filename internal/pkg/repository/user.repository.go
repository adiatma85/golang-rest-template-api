package repository

import (
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	"github.com/adiatma85/golang-rest-template-api/pkg/crypto"
	"github.com/adiatma85/golang-rest-template-api/pkg/helpers"
)

// Local variable
var (
	userRepository *UserRepository
)

// Contract of User Repository
type UserRepositoryInterface interface {
	Create(user models.User) (models.User, error)
	GetAll() (*[]models.User, error)
	Query(pagination helpers.Pagination) (*helpers.Pagination, error)
	QueryWithCondition(q *models.User, pagination helpers.Pagination) (*helpers.Pagination, error)
	GetByEmail(email string) (*models.User, error)
	GetById(userId string) (*models.User, error)
	Update(user *models.User) error
	Delete(user *models.User) error
	DeleteWithIds(ids []uint64) error
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
	err := Create(&user)
	// If error when transaction to database i.e duplicate email
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Func to get All User without Pagination
func (repo *UserRepository) GetAll() (*[]models.User, error) {
	var users []models.User
	err := Find(&models.User{}, &users, []string{"Product"}, "id asc")
	return &users, err
}

// Func to get Query of WHERE with pagination but
func (repo *UserRepository) Query(pagination helpers.Pagination) (*helpers.Pagination, error) {
	var users []models.User
	outputPagination, _ := Query(&models.User{}, &users, pagination, []string{"Product"})
	return outputPagination, nil
}

// Query with existed body from client
func (repo *UserRepository) QueryWithCondition(q *models.User, pagination helpers.Pagination) (*helpers.Pagination, error) {
	var users []models.User
	outputPagination, _ := Query(q, &users, pagination, []string{"Product"})
	return outputPagination, nil
}

// Func to get single user from email
func (repo *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.Email = email
	_, err := First(&where, &user, []string{"Product"})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Func to Get User By Id
func (repo *UserRepository) GetById(userId string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.ID = helpers.ConvertStringtoUint(userId)
	_, err := First(&where, &user, []string{"Product"})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Function to update user according to user schema defined
func (repo *UserRepository) Update(user *models.User) error {
	if user.Password != "" {
		passwordHelper := crypto.GetPasswordCryptoHelper()
		hashedPassword, err := passwordHelper.HashAndSalt([]byte(user.Password))
		user.Password = hashedPassword
		if err != nil {
			return err
		}
	}
	return Save(user)
}

// Delete User By Model defined in controller
func (repo *UserRepository) Delete(user *models.User) error {
	_, err := DeleteByModel(user)
	if err != nil {
		return err
	}
	return nil
}

// Delete User by multiple ids
func (repo *UserRepository) DeleteWithIds(ids []uint64) error {
	_, err := DeleteByIDS(models.User{}, ids)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"strconv"

	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/models"
	"github.com/adiatma85/go-tutorial-gorm/pkg/helpers"
)

// NEED TO FINALIZE PRODUCT REPOSITORY

// Local variable
var (
	productRepository *ProductRepository
)

// Struct to implements contract or interface
type ProductRepository struct{}

// Contract of Product Repository
type ProductRepositoryInterface interface {
	Create(product models.Product) (models.Product, error)
	GetAll() (*[]models.Product, error)
	Query(q *models.Product, pagination helpers.Pagination) (*helpers.Pagination, error)
	GetByEmail(email string) (*models.Product, error)
	GetById(productId string) (*models.Product, error)
	Update(product *models.Product) error
	Delete(product *models.Product) error
}

// Func to return Product Repository instance
func GetProductRepository() ProductRepositoryInterface {
	if productRepository == nil {
		productRepository = &ProductRepository{}
	}
	return productRepository
}

// Func to Create Product
func (repo *ProductRepository) Create(product models.Product) (models.Product, error) {
	if err != nil {
		return models.Product{}, err
	}
	err := Create(&product)
	// If error when transaction to database i.e duplicate email
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// Func to get All User without Pagination
func (repo *ProductRepository) GetAll() (*[]models.Product, error) {
	var users []models.Product
	err := Find(&models.Product{}, &users, []string{}, "id asc")
	return &users, err
}

// Func to get Query of WHERE with Pagination
func (repo *ProductRepository) Query(q *models.Product, pagination helpers.Pagination) (*helpers.Pagination, error) {
	var products []models.Product
	outputPagination, _ := Query(q, &products, pagination, []string{})
	return outputPagination, nil
}

// Func to get single user from email
func (repo *ProductRepository) GetByEmail(email string) (*models.Product, error) {
	var user models.Product
	where := models.Product{}
	_, err := First(&where, &user, []string{})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Func to Get User By Id
func (repo *ProductRepository) GetById(userId string) (*models.Product, error) {
	var user models.Product
	where := models.Product{}
	where.ID, _ = strconv.ParseUint(userId, 10, 64)
	_, err := First(&where, &user, []string{})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Function to update user according to user schema defined
func (repo *ProductRepository) Update(user *models.Product) error {
	return Save(user)
}

// Delete User By Model defined in controller
func (repo *ProductRepository) Delete(user *models.Product) error {
	_, err = DeleteByModel(user)
	if err != nil {
		return err
	}
	return nil
}

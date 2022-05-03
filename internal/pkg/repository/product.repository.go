package repository

import (
	"strconv"

	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	"github.com/adiatma85/golang-rest-template-api/pkg/helpers"
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
	Query(pagination helpers.Pagination) (*helpers.Pagination, error)
	QueryWithCondition(q *models.Product, pagination helpers.Pagination) (*helpers.Pagination, error)
	GetByEmail(email string) (*models.Product, error)
	GetById(productId string) (*models.Product, error)
	Update(product *models.Product) error
	Delete(product *models.Product) error
	DeleteWithIds(ids []uint64) error
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
	err := Find(&models.Product{}, &users, []string{"User"}, "id asc")
	return &users, err
}

// Func to get Query of WHERE with Pagination
func (repo *ProductRepository) Query(pagination helpers.Pagination) (*helpers.Pagination, error) {
	var products []models.Product
	outputPagination, _ := Query(&models.Product{}, &products, pagination, []string{"User"})
	return outputPagination, nil
}

// Query with existed body from client
func (repo *ProductRepository) QueryWithCondition(q *models.Product, pagination helpers.Pagination) (*helpers.Pagination, error) {
	var products []models.Product
	outputPagination, _ := Query(q, &products, pagination, []string{"User"})
	return outputPagination, nil
}

// Func to get single user from email
func (repo *ProductRepository) GetByEmail(email string) (*models.Product, error) {
	var user models.Product
	where := models.Product{}
	_, err := First(&where, &user, []string{"User"})
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
	_, err := First(&where, &user, []string{"User"})
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
	_, err := DeleteByModel(user)
	if err != nil {
		return err
	}
	return nil
}

// Delete User by multiple ids
func (repo *ProductRepository) DeleteWithIds(ids []uint64) error {
	_, err := DeleteByIDS(models.Product{}, ids)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
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

// Func to get All Product without Pagination
func (repo *ProductRepository) GetAll() (*[]models.Product, error) {
	var products []models.Product
	err := Find(&models.Product{}, &products, []string{"User"}, "id asc")
	return &products, err
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

// Func to get single product from email
func (repo *ProductRepository) GetByEmail(email string) (*models.Product, error) {
	var product models.Product
	where := models.Product{}
	_, err := First(&where, &product, []string{"User"})
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Func to Get Product By Id
func (repo *ProductRepository) GetById(productId string) (*models.Product, error) {
	var product models.Product
	where := models.Product{}
	where.ID = helpers.ConvertStringtoUint(productId)
	_, err := First(&where, &product, []string{"User"})
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Function to update product according to product schema defined
func (repo *ProductRepository) Update(product *models.Product) error {
	return Save(product)
}

// Delete Product By Model defined in controller
func (repo *ProductRepository) Delete(product *models.Product) error {
	_, err := DeleteByModel(product)
	if err != nil {
		return err
	}
	return nil
}

// Delete Product by multiple ids
func (repo *ProductRepository) DeleteWithIds(ids []uint64) error {
	_, err := DeleteByIDS(models.Product{}, ids)
	if err != nil {
		return err
	}
	return nil
}

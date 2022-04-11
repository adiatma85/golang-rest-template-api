package repository

import (
	"fmt"
	"testing"

	"github.com/adiatma85/golang-rest-template-api/internal/pkg/db"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/repository"
	"github.com/adiatma85/golang-rest-template-api/test"
	"github.com/stretchr/testify/suite"
)

// User repo needed because product can not be an orphan in relationship
var userRepo repository.UserRepositoryInterface

// Struct for Producy Repository Suite
type ProductRepositorySuite struct {
	// We need this to use the suite functionalities from testify
	suite.Suite
	productRepo repository.ProductRepositoryInterface
}

// Main Function for Test Suite
func TestProductRepository(t *testing.T) {
	suite.Run(t, new(ProductRepositorySuite))
	// Clean up after all testing
	defer test.TearDownHelper()
}

// Setup Suite
func (suite *ProductRepositorySuite) SetupSuite() {
	// Initialize configuration
	test.SetupInitialize("../../../.env")
	db.SetupTestingDb(test.Host, test.Username, test.Password, test.Port, test.Database)

	// User Repository here because product can not be an orphan
	userRepo = repository.GetUserRepository()

	// Product repo to test
	suite.productRepo = repository.GetProductRepository()

	// inserting dummy user
	for _, user := range users {
		userRepo.Create(user)
	}

	// inserting dummy product
	for _, product := range products {
		suite.productRepo.Create(product)
	}
}

// Create Product Repository
func (suite *ProductRepositorySuite) TestCreateProduct_Positive() {
	// Creating product
	createdProduct, err := suite.productRepo.Create(willBeProduct)

	// Equal assertion
	suite.Equal(willBeProduct.Name, createdProduct.Name, "both of the name from dummy data and existed product should have the same value")
	suite.Equal(willBeProduct.Price, createdProduct.Price, "both of the price from dummy data and existed product should have the same value")
	suite.Equal(willBeProduct.UserId, createdProduct.UserId, "both of the user id from dummy data and existed product should have the same value")
	suite.NoError(err, "should have no error when creating new product with this parameter")
}

// Get All Product Test Repository
func (suite *ProductRepositorySuite) TestGetAllProduct_Positive() {
	_, err := suite.productRepo.GetAll()
	suite.NoError(err, "should have no error when fetching products (bulk fetch)")
}

// Test Query User with default pagination
func (suite *ProductRepositorySuite) TestQueryProductsWithDefaultPagination() {
	pagination, err := suite.productRepo.Query(pagination0)
	suite.Equal(10, pagination.GetLimit(), "pagination limit should have value of 10")
	suite.Equal(1, pagination.GetPage(), "pagination page should have value of 1")
	suite.Equal("Id desc", pagination.GetSort(), "pagination sort should have value of 'Id desc'")
	suite.NoError(err, "should have no error when fetching products (pagination fetch)")
}

// Test Query Product with pre-defined pagination
func (suite *ProductRepositorySuite) TestQueryProductsWithPreDefinedPagination() {
	pagination, err := suite.productRepo.Query(pagination1)
	suite.Equal(pagination1.GetLimit(), pagination.GetLimit(), fmt.Sprintf("pagination limit should have value %d", pagination1.GetLimit()))
	suite.Equal(pagination1.GetPage(), pagination.GetPage(), fmt.Sprintf("pagination page should have value %d", pagination1.GetPage()))
	suite.Equal("Id desc", pagination.GetSort(), "pagination sort should have value of 'Id desc'")
	suite.NoError(err, "should have no error when fetching products (pagination fetch)")
}

// Test get product from id
func (suite *ProductRepositorySuite) TestGetByIdPositive() {
	product, err := suite.productRepo.GetById("1")
	suite.Equal(uint64(1), product.ID, "both of the id from client data and existed product should have the same value")
	suite.Equal(products[0].Name, product.Name, "both of the name from dummy data and existed product should have the same value")
	suite.Equal(products[0].Price, product.Price, "both of the price from dummy data and existed product should have the same value")
	suite.NoError(err, "should have no error when fetching product (singular fetch by id)")
}

// Test get User from id but non exist
func (suite *ProductRepositorySuite) TestGetByIdNegative() {
	_, err := suite.productRepo.GetById("1000")
	suite.Error(err, "should have an error when fetching product (singular fetch by id)")
	suite.Equal(err.Error(), "record not found")
}

// Update One Product Test Repository
func (suite *ProductRepositorySuite) TestUpdateAProduct_Positive() {
	// model id
	insideModelId := models.Model{
		ID: 1,
	}

	// definedProduct
	updateProduct := models.Product{
		Model: insideModelId,
		Name:  "changing name",
	}

	err := suite.productRepo.Update(&updateProduct)

	// Equal assertion to make sure that updated attribute is updated
	updatedProduct, _ := suite.productRepo.GetById("1")
	suite.Equal(updateProduct.ID, updatedProduct.ID, "both of the 'id' product from client and existed product should have the same value")
	suite.Equal(updateProduct.Name, updatedProduct.Name, "both of the 'name' product from client and existed product should have the same value")
	suite.NoError(err, "should have no error when updating product")
}

// Delete a Product Repository
func (suite *ProductRepositorySuite) TestDeleteAProduct_Positive() {
	product := models.Product{
		Model: models.Model{
			ID: 2,
		},
	}
	err := suite.productRepo.Delete(&product)
	suite.NoError(err, "should have no error when deleting product (singular delete by id)")
}

package unit

import (
	"testing"

	"github.com/adiatma85/golang-rest-template-api/internal/pkg/db"
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
	for _, model := range test.Models {
		db.GetDB().Migrator().DropTable(model)
	}
}

// Setup Suite
func (suite *ProductRepositorySuite) SetupSuite() {
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

}

// Get All Product Test Repository
func (suite *ProductRepositorySuite) TestGetAllProduct_Positive() {

}

// Test Query User with default pagination
func (suite *ProductRepositorySuite) QueryProductsWithDefaultPagination() {

}

// Test Query Product with pre-defined pagination
func (suite *ProductRepositorySuite) QueryProductsWithPreDefinedPagination() {

}

// Test get product from id
func (suite *ProductRepositorySuite) TestGetByIdPositive() {

}

// Test get User from id but non exist
func (suite *ProductRepositorySuite) TestGetByIdNegative() {

}

// Update One Product Test Repository
func (suite *ProductRepositorySuite) TestUpdateAProduct_Positive() {

}

// Delete a Product Repository
func (suite *ProductRepositorySuite) TestDeleteAProduct_Positive() {

}

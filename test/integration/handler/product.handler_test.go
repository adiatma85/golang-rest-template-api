package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/adiatma85/golang-rest-template-api/internal/api/handler"
	"github.com/adiatma85/golang-rest-template-api/test/mocks/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// Struct for Product Handler Suite
type ProductHandlerSuite struct {
	// we need this to use the suite functionalities from testify
	suite.Suite
	// the mocked version of the service
	repository *repository.ProductRepositoryInterface
	// the functionalities we need to test
	handler handler.ProductHandlerInterface
	// testing server to be used the handler
	testingServer *httptest.Server
}

// Main Function for Test Suite
func TestProductHandler(t *testing.T) {
	suite.Run(t, new(ProductHandlerSuite))
}

// Function to initialize the test suite
func (suite *ProductHandlerSuite) SetupSuite() {
	// create a mocked version of service
	repository := new(repository.ProductRepositoryInterface)

	productHandler := handler.GetProductHandler()

	// create default server using gin, then register all endpoints
	router := gin.Default()

	// create and run the testing server
	testingServer := httptest.NewServer(router)

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.testingServer = testingServer
	suite.repository = repository
	suite.handler = productHandler
}

// Wrrapping up after testing is finished
func (suite *ProductHandlerSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}

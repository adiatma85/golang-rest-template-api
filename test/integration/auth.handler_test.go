package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/adiatma85/golang-rest-template-api/internal/api/handler"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/validator"
	responseHelper "github.com/adiatma85/golang-rest-template-api/pkg/response"
	"github.com/adiatma85/golang-rest-template-api/test/mocks/helpers"
	"github.com/adiatma85/golang-rest-template-api/test/mocks/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// Struct for Product Handler Suite
type AuthHandlerSuite struct {
	// we need this to use the suite functionalities from testify
	suite.Suite
	// the mocked version of the service
	repository *repository.UserRepositoryInterface
	// the mocked version of the password crypto helper
	passwordHelper *helpers.PasswordCryptoHelper
	// the mocked version of the jwt crypto helper
	jwtHelper *helpers.JWTCryptoHelper
	// the functionalities we need to test
	handler handler.AuthHandlerInterface
	// testing server to be used the handler
	testingServer *httptest.Server
}

// Main Function for Test Suite
func TestAuthHandler(t *testing.T) {
	suite.Run(t, new(AuthHandlerSuite))
}

// Function to initialize the test suite
func (suite *AuthHandlerSuite) SetupSuite() {
	// create a mocked version of repository
	repository := new(repository.UserRepositoryInterface)

	// create a mocked version of password crypto helper
	passwordHelper := new(helpers.PasswordCryptoHelper)

	// create a mocked version of jwt crypto helper
	jwtHelper := new(helpers.JWTCryptoHelper)

	authHandler := handler.GetAuthHandler()

	// create default server using gin, then register all endpoints
	router := gin.Default()
	// NEED TO REGISTER THE ENDPOINTS IN HERE
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("login", authHandler.AuthLogin)
		authGroup.POST("register", authHandler.AuthRegister)
	}

	// create and run the testing server
	testingServer := httptest.NewServer(router)

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.testingServer = testingServer
	suite.passwordHelper = passwordHelper
	suite.jwtHelper = jwtHelper
	suite.repository = repository
	suite.handler = authHandler
}

// Wrrapping up after testing is finished
func (suite *AuthHandlerSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}

// Func to mock test Login Success
func (suite *AuthHandlerSuite) TestLoginPositive() {
	// an example login request for the test
	loginData := validator.LoginRequest{
		Email:    "admin@admin.com",
		Password: "password",
	}

	// example of existed user data
	existedUser := models.User{
		Model: models.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     "admin",
		Email:    "admin@admin.com",
		Password: "password",
		Product:  []models.Product{},
	}

	// example of token
	token := "random long token"

	// specify that inside handler's AuthLogin method
	// repository's GetByemail method will be called
	suite.repository.On("GetByEmail", loginData.Email).Return(existedUser, nil)

	// specify that inside handler's AuthLogin method
	// passwordHelper's ComparePassword method will be called
	suite.passwordHelper.On("ComparePassword", existedUser.Password, []byte(loginData.Password)).Return(true)

	// specify that inside handler's AuthLogin method
	// jwtHelper's GenerateToken method will be called
	suite.jwtHelper.On("GenerateToken", fmt.Sprint(existedUser.ID)).Return(token, nil)

	// marshalling and some assertion
	requestBody, err := json.Marshal(&loginData)
	suite.NoError(err, "can not marshal struct to json")

	// calling the testing server given the provided request body
	// NEED TO ADD base endpoint in here
	response, err := http.Post(fmt.Sprintf("%s/api/v1/auth/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "error when doing POST to login endpoints")
	defer response.Body.Close()

	// unmarshalling the response
	responseBody := responseHelper.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	// // running assertions to make sure that our method does the correct thing
	suite.Equal("success login", responseBody.Message)
	suite.repository.AssertExpectations(suite.T())
}

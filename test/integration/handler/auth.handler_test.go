package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/adiatma85/golang-rest-template-api/internal/api/handler"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/db"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	appRepository "github.com/adiatma85/golang-rest-template-api/internal/pkg/repository"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/validator"
	"github.com/adiatma85/golang-rest-template-api/pkg/crypto"
	appResponseHelper "github.com/adiatma85/golang-rest-template-api/pkg/response"
	"github.com/adiatma85/golang-rest-template-api/test"
	"github.com/adiatma85/golang-rest-template-api/test/mocks/helpers"
	"github.com/adiatma85/golang-rest-template-api/test/mocks/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// Struct for Product Handler Suite
type AuthHandlerSuite struct {
	suite.Suite
	repository     *repository.UserRepositoryInterface
	passwordHelper *helpers.PasswordCryptoHelper
	jwtHelper      *helpers.JWTCryptoHelper
	handler        handler.AuthHandlerInterface
	testingServer  *httptest.Server
}

// Main Function for Test Suite
func TestAuthHandler(t *testing.T) {
	suite.Run(t, new(AuthHandlerSuite))
	// Clean up after all testing
	defer test.TearDownHelper()
}

// Function to initialize the test suite
func (suite *AuthHandlerSuite) SetupSuite() {
	// Initialize database configuration
	test.SetupInitialize("../../../.env")
	db.SetupTestingDb(test.Host, test.Username, test.Password, test.Port, test.Database)

	// Create a mocked version of repository
	repository := new(repository.UserRepositoryInterface)

	// Create a mocked version of password crypto helper
	passwordHelper := new(helpers.PasswordCryptoHelper)

	// Create a mocked version of jwt crypto helper
	jwtHelper := new(helpers.JWTCryptoHelper)

	authHandler := handler.GetAuthHandler()

	// Create default server using gin, then register all endpoints
	router := gin.Default()
	// List of Endpoints that need to be tested
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("login", authHandler.AuthLogin)
		authGroup.POST("register", authHandler.AuthRegister)
	}

	// Create and run the testing server
	testingServer := httptest.NewServer(router)

	// Assign the dependencies we need as the suite properties
	// We need this to run the tests
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
		Email:    "admin2@admin.com",
		Password: "password",
	}

	// example of existed user data
	existedUser := models.User{
		Name:     "admin",
		Email:    "admin2@admin.com",
		Password: "password",
		Product:  []models.Product{},
	}

	var err error

	// Example of token
	token := "random long token"

	// Insert random user to database
	userRepository := appRepository.GetUserRepository()
	// Change the password to hash and salt
	passwordHelper := crypto.GetPasswordCryptoHelper()
	existedUser.Password, err = passwordHelper.HashAndSalt([]byte(existedUser.Password))
	suite.NoError(err, "can not change dummy password")

	// Insert to user repository
	_, err = userRepository.Create(existedUser)
	suite.NoError(err, "can not inserting dummy data for user")

	// Specify that inside handler's AuthLogin method
	// repository's GetByemail method will be called
	suite.repository.On("GetByEmail", loginData.Email).Return(existedUser, nil)

	// Specify that inside handler's AuthLogin method
	// passwordHelper's ComparePassword method will be called
	suite.passwordHelper.On("ComparePassword", existedUser.Password, []byte(loginData.Password)).Return(true)

	// Specify that inside handler's AuthLogin method
	// jwtHelper's GenerateToken method will be called
	suite.jwtHelper.On("GenerateToken", fmt.Sprint(existedUser.ID)).Return(token, nil)

	// Marshalling and some assertion
	requestBody, err := json.Marshal(&loginData)
	suite.NoError(err, "can not marshal struct to json")

	// Calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/api/v1/auth/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "error when doing POST to login endpoints")
	defer response.Body.Close()

	// unmarshalling the response
	responseBody := appResponseHelper.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	// running assertions to make sure that our method does the correct thing
	suite.Equal("success login", responseBody.Message)
}

func (suite *AuthHandlerSuite) TestRegisterPositive() {
	// an example register request for the test
	registerData := validator.RegisterRequest{
		Name:     "admin",
		Email:    "admin@admin.com",
		Password: "password",
	}

	// will be existed user data
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

	// Specify that inside handler's AuthLogin method
	// passwordHelper's ComparePassword method will be called
	suite.passwordHelper.On("HashAndSalt", []byte(registerData.Password)).Return("randomHashAndSalt")

	// Specify that inside handler's AuthLogin method
	// jwtHelper's GenerateToken method will be called
	suite.jwtHelper.On("GenerateToken", fmt.Sprint(existedUser.ID)).Return(token, nil)

	// Marshalling and some assertion
	requestBody, err := json.Marshal(&registerData)
	suite.NoError(err, "can not marshal struct to json")

	// Calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/api/v1/auth/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "error when doing POST to register endpoints")
	defer response.Body.Close()

	// Unmarshalling the response
	responseBody := appResponseHelper.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	// Running assertions to make sure that our method does the correct thing
	suite.Equal("success register new user", responseBody.Message)
}

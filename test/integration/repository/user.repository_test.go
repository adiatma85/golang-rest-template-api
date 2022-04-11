package repository

import (
	"fmt"
	"testing"

	"github.com/adiatma85/golang-rest-template-api/internal/pkg/db"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/repository"
	"github.com/adiatma85/golang-rest-template-api/pkg/crypto"
	"github.com/adiatma85/golang-rest-template-api/test"
	"github.com/stretchr/testify/suite"
)

// Struct for User Repository Suite
type UserRepositorySuite struct {
	// We need this to use the suite functionalities from testify
	suite.Suite
	userRepo repository.UserRepositoryInterface
}

// Main Function for Test Suite
func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
	// Clean up after all testing
	defer test.TearDownHelper()
}

// Function to initialize the test suite
func (suite *UserRepositorySuite) SetupSuite() {
	// Initialize configuration
	test.SetupInitialize("../../../.env")
	db.SetupTestingDb(test.Host, test.Username, test.Password, test.Port, test.Database)
	suite.userRepo = repository.GetUserRepository()

	// inserting dummy user
	for _, user := range users {
		suite.userRepo.Create(user)
	}
}

// Create User instance Test
func (suite *UserRepositorySuite) TestCreateUser_Positive() {
	// Creating user
	createdUser, err := suite.userRepo.Create(willBeUser)

	// Equal assertion
	suite.Equal(willBeUser.Name, createdUser.Name, "both of the name from dummy data and existed user should have the same value")
	suite.Equal(willBeUser.Email, createdUser.Email, "both of the email from dummy data and existed user should have the same value")
	// assume password hashed outside unit repository
	suite.Equal(willBeUser.Password, createdUser.Password, "both of the password from dummy data and existed user should have the same value")
	suite.NoError(err, "should have no error when creating new user with this parameter")
}

// Get All User instances Test
func (suite *UserRepositorySuite) TestGetAllUser_Positive() {
	_, err := suite.userRepo.GetAll()
	suite.NoError(err, "should have no error when fetching users (bulk fetch)")
}

// Test Query User with default pagination
func (suite *UserRepositorySuite) TestQueryUsersWithDefaultPagination() {
	pagination, err := suite.userRepo.Query(pagination0)
	// pagination dilengkapi
	suite.Equal(10, pagination.GetLimit(), "pagination limit should have value of 10")
	suite.Equal(1, pagination.GetPage(), "pagination page should have value of 1")
	suite.Equal("Id desc", pagination.GetSort(), "pagination sort should have value of 'Id desc'")
	suite.NoError(err, "should have no error when fetching users (pagination fetch)")
}

// Test Query User with pre-defined pagination
func (suite *UserRepositorySuite) TestQueryUsersWithPreDefinedPagination() {
	pagination, err := suite.userRepo.Query(pagination1)
	// pagination dilengkapi
	suite.Equal(pagination1.GetLimit(), pagination.GetLimit(), fmt.Sprintf("pagination limit should have value %d", pagination1.GetLimit()))
	suite.Equal(pagination1.GetPage(), pagination.GetPage(), fmt.Sprintf("pagination page should have value %d", pagination1.GetPage()))
	suite.Equal("Id desc", pagination.GetSort(), "pagination sort should have value of 'Id desc'")
	suite.NoError(err, "should have no error when fetching users (pagination fetch)")
}

// Test get user from id
func (suite *UserRepositorySuite) TestGetByIdPositive() {
	user, err := suite.userRepo.GetById("1")
	suite.Equal(uint64(1), user.ID, "both of the id from client data and existed user should have the same value")
	suite.Equal(users[0].Name, user.Name, "both of the name from dummy data and existed user should have the same value")
	suite.Equal(users[0].Email, user.Email, "both of the email from dummy data and existed user should have the same value")
	suite.NoError(err, "should have no error when fetching user (singular fetch by id)")
}

// Test get user from their email
func (suite *UserRepositorySuite) TestGetByEmailPositive() {
	existedUser, err := suite.userRepo.GetByEmail(willBeUser.Email)
	suite.Equal(willBeUser.Name, existedUser.Name, "both of the name from dummy data and existed user should have the same value")
	suite.Equal(willBeUser.Email, existedUser.Email, "both of the email from client and existed user should have the same value")
	suite.NoError(err, "should have no error when fetching user (singular fetch by id)")
}

// Test get User from id but non exist
func (suite *UserRepositorySuite) TestGetByIdNegative() {
	_, err := suite.userRepo.GetById("1000")
	suite.Error(err, "should have an error when fetching user (singular fetch by id)")
	suite.Equal(err.Error(), "record not found")
}

// Update One User Test Repository
func (suite *UserRepositorySuite) TestUpdateAUser_Positive() {

	// model id
	insideModelId := models.Model{
		ID: 1,
	}

	// definedUser
	updateUser := models.User{
		Model: insideModelId,
		Name:  "changing name",
	}
	err := suite.userRepo.Update(&updateUser)

	// Equal assertion to make sure that updated attribute is updated
	updatedUser, _ := suite.userRepo.GetById("1")
	suite.Equal(updateUser.ID, updatedUser.ID, "both of the 'id' user from client and existed user should have the same value")
	suite.Equal(updateUser.Name, updatedUser.Name, "both of the 'name' user from client and existed user should have the same value")
	suite.NoError(err, "should have no error when updating user (without change the password)")
}

// Update One User Test Repository with Password change
func (suite *UserRepositorySuite) TestUpdateAUserwithPassword_Positive() {
	// model id
	insideModelId := models.Model{
		ID: 1,
	}

	// definedUser
	updateUser := models.User{
		Model:    insideModelId,
		Password: "changing password",
	}
	err := suite.userRepo.Update(&updateUser)
	// Equal assertion to make sure that updated attribute is updated
	updatedUser, _ := suite.userRepo.GetById("1")
	passwordHelper := crypto.GetPasswordCryptoHelper()
	condition := passwordHelper.ComparePassword(updatedUser.Password, []byte("changing password"))
	suite.Equal(updateUser.ID, updatedUser.ID, "both of the 'id' user from client and existed user should have the same value")
	suite.True(condition, "should return true when between hashed Password and plain password")
	suite.NoError(err, "should have no error when updating user (without change the password)")
}

// Delete A User Test Repository
func (suite *UserRepositorySuite) TestDeleteAUser_Positive() {
	user := models.User{
		Model: models.Model{
			ID: 2,
		},
	}
	err := suite.userRepo.Delete(&user)
	suite.NoError(err, "should have no error when deleting user (singular delete by id)")
}

package unit

import (
	"testing"

	"github.com/adiatma85/golang-rest-template-api/internal/pkg/db"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/repository"
	"github.com/adiatma85/golang-rest-template-api/pkg/crypto"
	"github.com/adiatma85/golang-rest-template-api/pkg/helpers"
	"github.com/adiatma85/golang-rest-template-api/test"
	"github.com/stretchr/testify/suite"
)

type UserRepositorySuite struct {
	// We need this to use the suite functionalities from testify
	suite.Suite
	userRepo repository.UserRepositoryInterface
}

// Function to initialize the test suite
func (suite *UserRepositorySuite) SetupSuite() {
	// Initialize configuration
	db.SetupTestingDb(test.Host, test.Username, test.Password, test.Port, test.Database)
	suite.userRepo = repository.GetUserRepository()

	// inserting dummy user
	for _, user := range users {
		suite.userRepo.Create(user)
	}
}

// Create User Test Repository
func (suite *UserRepositorySuite) TestCreateUser_Positive() {
	userPassword := "Password"
	passwordHelper := crypto.GetPasswordCryptoHelper()
	hashedPassword, _ := passwordHelper.HashAndSalt([]byte(userPassword))

	// Models from handler should like this
	user := models.User{
		Name:     "Korenia",
		Password: hashedPassword,
		Email:    "korenia@example.com",
	}

	createdUser, err := suite.userRepo.Create(user)

	// Equal assertion
	suite.Equal(user.Name, createdUser.Name)
	suite.Equal(user.Email, createdUser.Email)
	suite.Equal(user.Password, createdUser.Password)
	suite.NoError(err, "no error when creating new user")
}

// // Get All User Test Repository
func (suite *UserRepositorySuite) TestGetAllUser_Positive() {
	_, err := suite.userRepo.GetAll()
	suite.NoError(err, "no error when fetching all user")
}

// // Test Query User with default pagination
func (suite *UserRepositorySuite) TestQueryUsersWithDefaultPagination() {
	pagination, err := suite.userRepo.Query(helpers.Pagination{})
	suite.Equal(pagination.GetLimit(), 10)
	suite.Equal(pagination.GetPage(), 1)
	suite.NoError(err, "no error when pagination users")
}

// Test Query User with pre-defined pagination
func (suite *UserRepositorySuite) TestQueryUsersWithPreDefinedPagination() {
	definedPagination := helpers.Pagination{
		Limit: 5,
		Page:  1,
	}
	pagination, err := suite.userRepo.Query(definedPagination)
	suite.Equal(pagination.Limit, 5)
	suite.Equal(pagination.Page, 1)
	suite.NoError(err, "no error when pagination users")
}

// Test get user from id
func (suite *UserRepositorySuite) TestGetByIdPositive() {
	_, err := suite.userRepo.GetById("1")
	suite.NoError(err, "no error when fetching specific user")
}

// Test get user from their email
func (suite *UserRepositorySuite) TestGetByEmailPositive() {
	searchedEmail := "korenia@example.com"
	existedUser, err := suite.userRepo.GetByEmail(searchedEmail)
	suite.Equal(searchedEmail, existedUser.Email, "the email should same")
	suite.NoError(err, "error happen when fetch single user by email")
}

// Test get User from id but non exist
func (suite *UserRepositorySuite) TestGetByIdNegative() {
	_, err := suite.userRepo.GetById("1000")
	suite.Error(err, "should have been error when search non-existent resources")
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
	suite.Equal(updateUser.Name, updatedUser.Name)

	suite.NoError(err, "no error when updating particular user")
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
	suite.True(condition, "updated password is match")
	suite.NoError(err, "no error when updating particular user")
}

// Delete A User Test Repository
func (suite *UserRepositorySuite) TestDeleteAUser_Positive() {
	user := models.User{
		Model: models.Model{
			ID: 2,
		},
	}
	err := suite.userRepo.Delete(&user)
	suite.NoError(err, "no error when deleting")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
	// Clean up after all testing
	for _, model := range test.Models {
		db.GetDB().Migrator().DropTable(model)
	}
}

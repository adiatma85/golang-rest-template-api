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

// Func to teardown after testing
func (suite *UserRepositorySuite) TearDownTest() {
	for _, model := range test.Models {
		db.GetDB().Migrator().DropTable(model)
	}
}

// Create User Test Repository
func (suite *UserRepositorySuite) CreateUser_Positive() {
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

// Get All User Test Repository
func (suite *UserRepositorySuite) GetAllUser_Positive() {
	_, err := suite.userRepo.GetAll()
	suite.NoError(err, "no error when fetching all user")
}

// Test Query User with default pagination
func (suite *UserRepositorySuite) QueryUsersWithDefaultPagination() {
	pagination, err := suite.userRepo.Query(helpers.Pagination{})
	suite.Equal(pagination.Limit, 10)
	suite.Equal(pagination.Page, 1)
	suite.NoError(err, "no error when pagination users")
}

// Test Query User with pre-defined pagination
func (suite *UserRepositorySuite) QueryUsersWithPreDefinedPagination() {
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
func (suite *UserRepositorySuite) GetByIdPositive() {
	_, err := suite.userRepo.GetById("1")
	suite.NoError(err, "no error when fetching specific user")
}

// Test get User from id but non exist
func (suite *UserRepositorySuite) GetByIdNegative() {
	_, err := suite.userRepo.GetById("-1")
	suite.Error(err, "error when fetching non existent id")
}

// Update One User Test Repository
func (suite *UserRepositorySuite) UpdateAUser_Positive() {

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
func (suite *UserRepositorySuite) UpdateAUserwithPassword_Positive() {
	// definedUser
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
	condition := passwordHelper.ComparePassword(updatedUser.Password, []byte(updateUser.Password))
	suite.True(condition, "updated password is match")
	suite.NoError(err, "no error when updating particular user")
}

// Update One User Test Repository with non-existet id
func (suite *UserRepositorySuite) UpdateAUser_Negative() {
	// non existent user
	updateUser := models.User{
		Model: models.Model{
			ID: 1000, // the number is so high that it exceed the testing case
		},
		Name: "changing name",
	}
	err := suite.userRepo.Update(&updateUser)
	suite.Error(err, "error when update a user")
}

// Delete A User Test Repository
func (suite *UserRepositorySuite) DeleteAUser_Positive() {
	user := models.User{
		Model: models.Model{
			ID: 1,
		},
	}
	err := suite.userRepo.Delete(&user)
	suite.NoError(err, "no error when deleting")
}

// Delete a User Test with non-existend id
func (suite *UserRepositorySuite) DeleteAUser_Negative() {
	// failed test here
	user := models.User{
		Model: models.Model{
			ID: 1000, // the number is so high that it exceed the testing case
		},
	}
	err := suite.userRepo.Delete(&user)
	suite.Error(err, "error when deleting user or users with specific condition")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

package unit

import (
	"testing"

	"github.com/adiatma85/golang-rest-template-api/internal/pkg/db"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/repository"
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
}

// Func to teardown after testing
func (suite *UserRepositorySuite) TearDownTest() {
	for _, model := range test.Models {
		db.GetDB().Migrator().DropTable(model)
	}
}

// Create User Test Repository
func (suite *UserRepositorySuite) TestCreateUser_Positive() {
	user := models.User{
		Name:     "Ramdani Koernia",
		Password: "Password",
		Email:    "adiatma85@gmail.com",
	}

	_, err := suite.userRepo.Create(user)
	suite.NoError(err, "no error when creating new user")
}

// Get All User Test Repository
func (suite *UserRepositorySuite) TestGetAllUser_Positive() {

}

// Update One User Test Repository
func (suite *UserRepositorySuite) UpdateAUser_Positive() {

}

// Delete A User Test Repository
func (suite *UserRepositorySuite) DeleteAUser_Positive() {

}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

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

// Func Setup for Suite
func (suite *UserRepositorySuite) SetupSuite() {
	// Function to initialize the test suite
	// Initialize configuration
	db.SetupTestingDb(test.Host, test.Username, test.Password, test.Port, test.Database)
	suite.userRepo = repository.GetUserRepository()
}

func (suite *UserRepositorySuite) TearDownTest() {
	// fmt.Println("testing is finished")
}

func (suite *UserRepositorySuite) TestCreateUser_Positive() {
	user := models.User{
		Name:     "Ramdani Koernia",
		Password: "Password",
		Email:    "adiatma85@gmail.com",
	}

	_, err := suite.userRepo.Create(user)
	suite.NoError(err, "no error when creating new user")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

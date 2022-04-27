package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User Repository Suite
type UserRepositorySuite struct {
	suite.Suite
	Db             *gorm.DB
	Mock           sqlmock.Sqlmock
	UserRepository repository.UserRepositoryInterface
	User           *models.User
}

// Main Function for Test Suite
func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

// Will be Random Generated User for testing
var randomGeneratedUser = &models.User{
	Model: models.Model{
		ID: 1,
	},
	Name:     "Korenia",
	Email:    "korenia@example.com",
	Password: "Password",
	Product:  []models.Product{},
}

// Setup Suite
func (suite *UserRepositorySuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, suite.Mock, err = sqlmock.New()
	require.NoError(suite.T(), err)

	suite.Db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	require.NoError(suite.T(), err)
}

// After Test
func (suite *UserRepositorySuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.Mock.ExpectationsWereMet())
}

// Function to test Get All Function in user repository
func (suite *UserRepositorySuite) TestUserRepositoryGetAll() {
	query := `SELECT id, name, email FROM users`

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(randomGeneratedUser.ID, randomGeneratedUser.Name, randomGeneratedUser.Email, randomGeneratedUser.Password)

	suite.Mock.ExpectQuery(query).WillReturnRows(rows)
	userRepository := repository.GetUserRepository()
	users, err := userRepository.GetAll()
	assert.NotEmpty(suite.T(), users)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), users, 1)
}

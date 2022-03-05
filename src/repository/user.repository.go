package repository

import (
	"github.com/adiatma85/go-tutorial-gorm/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserById(userId string) (models.User, error)
	InsertUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	// VerifyCredential(email, password string) interface{}
	// IsDuplicateEmail(email string) *gorm.DB
	FindUserByEmail(email string) (models.User, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return &userRepository{
		connection: connection,
	}
}

func (ur *userRepository) FindUserById(userId string) (models.User, error) {
	var user models.User
	res := ur.connection.Where("id = ?", userId).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (ur *userRepository) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	res := ur.connection.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (ur *userRepository) InsertUser(user models.User) (models.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	ur.connection.Save(&user)
	return user, nil
}

func (ur *userRepository) UpdateUser(user models.User) (models.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser models.User
		ur.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	ur.connection.Save(&user)
	return user, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

package service

import (
	"errors"

	"github.com/adiatma85/go-tutorial-gorm/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email, password string) error
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) VerifyCredential(email, password string) error {
	user, err := s.userRepo.FindUserByEmail(email)

	if err != nil {
		return err
	}

	isValidPassword := comparePassword(password, user.Password)

	if !isValidPassword {
		return errors.New("credential doesn't watch")
	}

	return nil
}

func comparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

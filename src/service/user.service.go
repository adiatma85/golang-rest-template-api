package service

import (
	"errors"
	"log"

	"github.com/adiatma85/go-tutorial-gorm/src/models"
	"github.com/adiatma85/go-tutorial-gorm/src/repository"
	_user "github.com/adiatma85/go-tutorial-gorm/src/service/user"
	"github.com/adiatma85/go-tutorial-gorm/src/validator"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(registerValidator validator.RegisterValidator) (*_user.UserResponse, error)
	UpdateUser(updateValidator validator.UserUpdateValidator) (*_user.UserResponse, error)
	FindUserByEmail(email string) (*_user.UserResponse, error)
	FindUserById(id string) (*_user.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (us *userService) CreateUser(registerRequest validator.RegisterValidator) (*_user.UserResponse, error) {
	user, err := us.userRepo.FindUserByEmail(registerRequest.Email)

	if err == nil {
		return nil, errors.New("user already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = smapping.FillStruct(&user, smapping.MapFields(&registerRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	user, _ = us.userRepo.InsertUser(user)
	res := _user.NewUserResponse(user)
	return &res, nil
}

func (us *userService) UpdateUser(updateRequest validator.UserUpdateValidator) (*_user.UserResponse, error) {
	user := models.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&updateRequest))

	if err != nil {
		return nil, err
	}

	user, err = us.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	res := _user.NewUserResponse(user)
	return &res, nil
}

func (us *userService) FindUserByEmail(email string) (*_user.UserResponse, error) {
	user, err := us.userRepo.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	userResponse := _user.NewUserResponse(user)
	return &userResponse, nil
}

func (us *userService) FindUserById(id string) (*_user.UserResponse, error) {
	user, err := us.userRepo.FindUserById(id)

	if err != nil {
		return nil, err
	}

	userResponse := _user.UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}

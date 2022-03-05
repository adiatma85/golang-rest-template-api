package service

import "github.com/adiatma85/go-tutorial-gorm/src/repository"

type UserService interface {
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

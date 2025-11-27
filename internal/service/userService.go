package service

import (
	"errors"
	"go-microservices/internal/domain"
	"go-microservices/internal/dto"
	"go-microservices/internal/helper"
	"go-microservices/internal/repository"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email)

	log.Println(user)

	return &user, err
}

func (s UserService) Register(input dto.UserRegister) (string, error) {
	hPassword, err := s.Auth.CreateHashedPassword(input.Password)

	if err != nil {
		return "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.Phone,
	})

	if err != nil {
		return "", errors.New("could not create this user")
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.findUserByEmail(email)

	if err != nil {
		return "", errors.New("user does not exist")
	}

	err = s.Auth.VerifyPassword(password, user.Password)

	if err != nil {
		log.Println("password not matching")
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	//logic
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	//logic
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	//logic
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	//logic
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	//logic
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	//logic
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	//logic
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	//logic
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	//logic
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	//logic
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) (interface{}, error) {
	//logic
	return nil, nil
}

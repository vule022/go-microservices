package service

import (
	"errors"
	"fmt"
	"go-microservices/internal/domain"
	"go-microservices/internal/dto"
	"go-microservices/internal/repository"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email)

	log.Println(user)

	return &user, err
}

func (s UserService) Register(input dto.UserRegister) (string, error) {
	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	log.Println(user)
	userInfo := fmt.Sprintf("%v, %v, %v", user.ID, user.Email, user.UserType)

	return userInfo, err
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.findUserByEmail(email)

	if err != nil {
		return "", errors.New("user does not exist!")
	}

	return user.Email, nil
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

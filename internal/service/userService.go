package service

import (
	"errors"
	"fmt"
	"go-microservices/config"
	"go-microservices/internal/domain"
	"go-microservices/internal/dto"
	"go-microservices/internal/helper"
	"go-microservices/internal/repository"
	"go-microservices/pkg/notification"
	"log"
	"time"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config config.AppConfig
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

func (s UserService) isVerifiedUser(id uint) bool {
	currentUser, err := s.Repo.FindUserById(id)

	return err == nil && currentUser.Verified
}

func (s UserService) GetVerificationCode(e domain.User) error {
	if s.isVerifiedUser(e.ID) {
		return errors.New("user already verfied")
	}

	code, err := s.Auth.GenerateCode()

	if err != nil {
		return err
	}

	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = s.Repo.UpdateUser(e.ID, user)

	if err != nil {
		return errors.New("unable to update verification code")
	}

	user, _ = s.Repo.FindUserById(e.ID)

	notificationClient := notification.NewNotificationClient(s.Config)

	msg := fmt.Sprintf("Your verification code is: %v", code)

	err = notificationClient.SendSMS(user.Phone, msg)

	if err != nil {
		return errors.New("verification code not sent")
	}

	return nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	if s.isVerifiedUser(id) {
		return errors.New("user already verfied")
	}

	user, err := s.Repo.FindUserById(id)

	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code does not match")
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("code expired")

	}

	updateUser := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(id, updateUser)

	if err != nil {
		return errors.New("unable to verify")
	}

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

func (s UserService) BecomeSeller(id uint, input dto.SellerInput) (string, error) {
	user, _ := s.Repo.FindUserById(id)

	if user.UserType == domain.SELLER {
		return "", errors.New("you are already a seller")
	}

	seller, err := s.Repo.UpdateUser(id, domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.PhoneNumber,
		UserType:  domain.SELLER,
	})

	if err != nil {
		return "", err
	}

	token, err := s.Auth.GenerateToken(user.ID, user.Email, seller.UserType)

	err = s.Repo.CreateBankAccount(domain.BankAccount{
		BankAccountNumber: input.BankAccountNumber,
		SwiftCode:         input.SwiftCode,
		PaymentType:       input.PaymentType,
		UserId:            id,
	})

	return token, err
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

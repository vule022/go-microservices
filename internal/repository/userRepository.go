package repository

import (
	"errors"
	"go-microservices/internal/domain"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)

	CreateBankAccount(e domain.BankAccount) error
}

type userRepository struct {
	db *gorm.DB
}

// CreateBankAccount implements UserRepository.
func (r *userRepository) CreateBankAccount(e domain.BankAccount) error {
	return r.db.Create(&e).Error
}

func (r userRepository) CreateUser(u domain.User) (domain.User, error) {
	err := r.db.Create(&u).Error

	if err != nil {
		log.Printf("creation error %v\n", err)
		return domain.User{}, errors.New("failed to create new user")
	}

	return u, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "email=?", email).Error

	if err != nil {
		log.Printf("error looking for user %v\n", err)
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error

	if err != nil {
		log.Printf("error looking for user %v\n", err)
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error

	if err != nil {
		log.Printf("error updating %v\n", err)
		return domain.User{}, errors.New("user not updated")
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

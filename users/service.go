package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	users := User{}
	users.Name = input.Name
	users.Email = input.Email
	users.Occupations = input.Occupations
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return users, err
	}
	users.PasswordHash = string(passwordHash)
	users.Roles = "USER"
	newUsers, err := s.repository.Save(users)
	if err != nil {
		return newUsers, err
	}
	return newUsers, nil
}

func (s *service) Login(input LoginInput) (User, error) {

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No users found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

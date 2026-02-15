package userservice

import (
	"fmt"
	"game-app/entity"
	"game-app/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(u entity.User) (entity.User, error)
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
}
type RegisterResponse struct {
	User entity.User
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	//validate phone number
	if !phonenumber.IsPhoneNumberValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number not valid")
	}

	//check uniqueness phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, err
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	//validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name length short")
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}
	// create new user in database or storage
	createdUser, err := s.repo.RegisterUser(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("not created user %w", err)
	}

	//return created user
	return RegisterResponse{
		User: createdUser,
	}, nil

}

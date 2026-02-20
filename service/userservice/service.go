package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"game-app/entity"
	"game-app/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
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

	//validate pass
	if len(req.Password) < 6 {
		return RegisterResponse{}, fmt.Errorf("password length short")
	}

	//TODO hash the password
	//pass := []byte(req.Password)
	//bcrypt.GenerateFromPassword(pass, 0)

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    GetMD5Hash(req.Password),
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

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type LoginResponse struct {
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	//check the exist phone number

	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password is not correct")
	}

	if user.Password != GetMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("username or password is not correct")
	}

	return LoginResponse{}, nil
	//check the pass

	//return ok

	panic("implement me")
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

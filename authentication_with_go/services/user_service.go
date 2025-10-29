package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error)
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserByID(id string) (*models.User, error) {
	fmt.Println("Getting user by ID in user service")
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error) {
	fmt.Println("Creating user in user service")

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return nil, err
	}

	user, err := u.userRepository.Create(payload.Username, payload.Email, hashedPassword)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	// Hardcoding for email and password
	email := payload.Email
	password := payload.Password

	// Step 1: Get the user by email from user repository
	user, err := u.userRepository.GetByEmail(email)

	// Step 2: If user does not exist, return error
	if err!= nil || user==nil {
		fmt.Println("User not found from the repository")
		if err!= nil {
			fmt.Println("Error thrown is: ", err)
			return "", err
		}
		return "", fmt.Errorf("user not found from the repository: %s", email)
	}

	// Step 3: If user exists, check password using utils
	passwordMatch := utils.CheckPassword(password, user.Password)
	if !passwordMatch {
		fmt.Println("The provided password does not match")
		return "", fmt.Errorf("the provided password does not match")
	}

	// Step 4: If password matches, return the jwt token, else return the error
	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
		"nbf":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	}
	token, err := utils.CreateJWTToken(user.ID, &jwtPayload)
	if err!= nil {
		fmt.Println("The token could not be created")
		return "", fmt.Errorf("the token could not be created: %v", err)
	}
	return token, nil
}
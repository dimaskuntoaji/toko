package services

import (
	"errors"
	"toko1/helper"
	"toko1/entity"
	"toko1/repository"
)

type UserService interface {
	Register(input entity.UserRegisterInput) (entity.UserResponseRegister, error)
	Login(input entity.UserLoginInput) (string, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

//# REGISTER USER

// Register godoc
// @Summary      Register account
// @Description  Register an account
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request body entity.UserRegisterInput true "Payload Body [RAW]"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users/register [post]
func (s *userService) Register(input entity.UserRegisterInput) (entity.UserResponseRegister, error) {
	var (
		user         entity.User
		userResponse entity.UserResponseRegister
	)
	user, _ = s.repository.GetUserByEmail(input.Email)
	if user.ID > 0 {
		return userResponse, errors.New("email already existed")
	}

	// Hash Password 
	password, err := helper.HashPassword(input.Password)
	if err != nil {
		return userResponse, errors.New("something wrong with password")
	}

	user.UserName = input.UserName
	user.Email = input.Email
	user.Password = password

	

	user, err = s.repository.CreateUser(user)

	userResponse.ID = user.ID
	userResponse.UserName = user.UserName
	userResponse.Email = user.Email
	userResponse.Password = user.Password
	userResponse.CreatedAt = user.CreatedAt

	return userResponse, helper.ReturnIfError(err)
}


//#LOGIN USER

// Login godoc
// @Summary      Login account
// @Description  Login an account
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request body entity.UserLoginInput true "Payload Body [RAW]"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users/login [post]
func (s *userService) Login(input entity.UserLoginInput) (string, error) {
	var token string

	user, _ := s.repository.GetUserByEmail(input.Email)
	if user.ID == 0 {
		return "", errors.New("user is not existed")
	}

	ok := helper.ComparePassword(user.Password, input.Password)
	if !ok {
		return token, errors.New("password is wrong")
	}

	token, err := helper.GenerateToken(user)
	return token, helper.ReturnIfError(err)

}
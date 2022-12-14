package services

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"tokobelanja/helpers"
	"tokobelanja/models"
	"tokobelanja/params"
	"tokobelanja/repositories"
)

type UserService interface {
	Register(userRegisterRequest params.UserRegisterRequest) (*params.UserRegisterResponse, error)
	Login(userLoginRequest params.UserLoginRequest) (*params.UserLoginResponse, error)
	TopUp(id uint, userTopUpRequest params.UserTopUpRequest) (*models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func (u userService) Register(userRegisterRequest params.UserRegisterRequest) (*params.UserRegisterResponse, error) {
	pwHash, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return &params.UserRegisterResponse{}, err
	}

	newUser := models.User{
		FullName: userRegisterRequest.FullName,
		Email:    userRegisterRequest.Email,
		Password: string(pwHash),
		Role:     "customer",
		Balance:  0,
	}

	_, err = u.repository.CreateUser(&newUser)

	if err != nil {
		return &params.UserRegisterResponse{}, err
	}
	resp := params.ParseToCreateUserResponse(&newUser)

	return &resp, nil
}

func (u userService) Login(userLoginRequest params.UserLoginRequest) (*params.UserLoginResponse, error) {
	userFound, err := u.repository.FindUserByEmail(userLoginRequest.Email)
	if err != nil {
		return &params.UserLoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userLoginRequest.Password))
	if err != nil {
		return &params.UserLoginResponse{}, err
	}

	token := helpers.GenerateToken(userFound.ID, userFound.Email, userFound.Role)

	resp := params.UserLoginResponse{}
	resp.Token = token

	return &resp, nil
}

func (u userService) TopUp(id uint, userTopUpRequest params.UserTopUpRequest) (*models.User, error) {
	userModel, err := u.repository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	userModel.Balance += userTopUpRequest.Balance
	userModel.UpdatedAt = time.Now()
	user, err := u.repository.UpdateUser(id, userModel)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

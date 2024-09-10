package usecase

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working with different use cases

type UserUseCase struct {
	UserRepository *repository.UserRepository
	Timeout        time.Duration
}

// method for working with the user repository
func NewUserUseCase(time time.Duration, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
		Timeout:        time,
	}
}

// method for registering user into the database
func (uc *UserUseCase) RegisterUser(c context.Context, user *domain.User) (*domain.UserResponse, error) {
	_, err := uc.UserRepository.GetByUserEmail(c, user.Email)
	if err == nil {
		return nil, errors.New("already registered user")
	}

	_, err = uc.UserRepository.GetByUserName(c, user.UserName)
	if err == nil {
		return nil, errors.New("already used username")
	}

	user.Id = primitive.NewObjectID()
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashed)
	user.Role = "USER"
	return uc.UserRepository.PostUser(c, user)
}

// method for updating user information
func (uc *UserUseCase) UpdateInfo(c context.Context, id string, userRequest *domain.UserRequest) (*domain.UserResponse, error) {
	return uc.UserRepository.UpdateUser(c, id, userRequest)
}

// method for updating user preference
func (uc *UserUseCase) UpdatePreference(c context.Context, id string, userPrefrence *domain.UserPreference) (*domain.UserResponse, error) {
	return uc.UserRepository.UpdatePreference(c, id, userPrefrence)
}

// method for getting user by using username
func (uc *UserUseCase) GetByUserName(c context.Context, username string) (*domain.UserResponse, error) {
	return uc.UserRepository.GetByUserName(c, username)
}

// method for getting user by using username
func (uc *UserUseCase) GetByUserNameForPass(c context.Context, username string) (*domain.User, error) {
	return uc.UserRepository.GetByUserNameForPass(c, username)
}

// method for logging user into the system
func (uc *UserUseCase) LoginUser(c context.Context, loginInfo *domain.LoginRequest) (*domain.User, error) {
	user, err := uc.UserRepository.GetByUserNameForPass(c, loginInfo.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		return nil, errors.New("invalid password or username")
	}
	return user, nil

}

// method for working with password update
func (uc *UserUseCase) UpdatePassword(c context.Context, username string, passInfo *domain.ChangePassword) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(passInfo.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	passInfo.NewPassword=string(hashed)
	return uc.UserRepository.UpdatePassword(c, username,passInfo)
}

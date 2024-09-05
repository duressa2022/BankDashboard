package usecase

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working with different use cases

type UserUseCase struct {
	userRepository *repository.UserRepository
}

// method for working with the user repository
func (uc *UserUseCase) NewUserUseCase(userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

// method for registering user into the database
func (uc *UserUseCase) RegisterUser(c context.Context, user *domain.User) (*domain.UserResponse, error) {
	_, err := uc.userRepository.GetByUserEmail(c, user.Email)
	if err == nil {
		return nil, errors.New("already registered user")
	}
	user.Id = primitive.NewObjectID()
	//hash the password here

	return uc.userRepository.PostUser(c, user)
}

// method for updating user information
func (uc *UserUseCase) UpdateInfo(c context.Context,id string,userRequest *domain.UserRequest)(*domain.UserResponse,error){
	return uc.userRepository.UpdateUser(c,id,userRequest)
}

// method for updating user preference 
func(uc *UserUseCase) UpdatePreference(c context.Context,id string,userPrefrence *domain.UserPreference)(*domain.UserPreference,error){
	return uc.userRepository.UpdatePreference(c,id,userPrefrence)
}

// method for getting user by using username
func (uc *UserUseCase) GetByUserName(c context.Context,username string)(*domain.UserResponse,error){
	return uc.userRepository.GetByUserName(c,username)
}

// method for logging user into the system
func (uc *UserUseCase) LoginUser(c context.Context,loginInfo *domain.LoginRequest)(string,string,error){
	_,err:=uc.userRepository.GetByUserName(c,loginInfo.Username)
	if err!=nil{
		return "","",err
	}
	//compare the password here
	//generate access and refresh token here
	access:=""
	refresh:=""
	return access,refresh,nil
}







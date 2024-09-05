package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type for working userinformation
type User struct {
	Id               primitive.ObjectID     `json:"id" bson:"_id"`
	Email            string                 `json:"email" bson:"email"`
	DateOfBirth      time.Time              `json:"dateOfBirth" bson:"dateOfBirth"`
	PermanentAddress string                 `json:"permanentAddress" bson:"permanentAddress"`
	PostalCode       string                 `json:"postalCode" bson:"postalCode"`
	UserName         string                 `json:"username" bson:"username"`
	Password         string                 `json:"password" bson:"password"`
	PresentAddress   string                 `json:"presentAddress" bson:"presentAddress"`
	City             string                 `json:"city" bson:"city"`
	Country          string                 `json:"country" bson:"country"`
	ProfilePicture   string                 `json:"profilePicture" bson:"profilePicture"`
	AccountBalance   float64                `json:"accountBalance" bson:"accountBalance"`
	Role             string                 `json:"role" bson:"role"`
	Preference       map[string]interface{} `preference:"country" bson:"preference"`
}

// type for working with user response
type UserResponse struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Name             string             `json:"name" bson:"name"`
	Email            string             `json:"email" bson:"email"`
	DateOfBirth      time.Time          `json:"dateOfBirth" bson:"dateOfBirth"`
	PermanentAddress string             `json:"permanentAddress" bson:"permanentAddress"`
	PostalCode       string             `json:"postalCode" bson:"postalCode"`
	UserName         string             `json:"username" bson:"username"`
	PresentAddress   string             `json:"presentAddress" bson:"presentAddress"`
	City             string             `json:"city" bson:"city"`
	Country          string             `json:"country" bson:"country"`
	ProfilePicture   string             `json:"profilePicture" bson:"profilePicture"`
	Preference       UserPreference     `preference:"country" bson:"preference"`
}

// type for working with user request
type UserRequest struct {
	Name             string    `json:"name" bson:"name"`
	Email            string    `json:"email" bson:"email"`
	DateOfBirth      time.Time `json:"dateOfBirth" bson:"dateOfBirth"`
	PermanentAddress string    `json:"permanentAddress" bson:"permanentAddress"`
	PostalCode       string    `json:"postalCode" bson:"postalCode"`
	UserName         string    `json:"username" bson:"username"`
	PresentAddress   string    `json:"presentAddress" bson:"presentAddress"`
	City             string    `json:"city" bson:"city"`
	Country          string    `json:"country" bson:"country"`
	ProfilePicture   string    `json:"profilePicture" bson:"profilePicture"`
}

// type for working user preference
type UserPreference struct {
	Currency                     string `json:"currency" bson:"currency"`
	SentOrReceiveDigitalCurrency bool   `json:"sentOrReceiveDigitalCurrency" bson:"sentOrReceiveDigitalCurrency"`
	ReceiveMerchantOrder          bool   `json:"receiveMerchantOrder" bson:"receiveMerchantOrder"`
	AccountRecommendations       bool   `json:"accountRecommendations" bson:"accountRecommendations"`
	TimeZone                     bool   `json:"timeZone" bson:"timeZone"`
	TwoFactorAuthentication      bool   `json:"twoFactorAuthentication" bson:"twoFactorAuthentication"`
}

// interface for working with user repo

type UserInterface interface {
	PostUser(c context.Context,userRequest *UserRequest)(*UserResponse,error)
	Update(c context.Context,userRequest *UserRequest)(*UserResponse,error)
	UpdatePreference(c context.Context,userPreference *UserPreference)(*UserPreference,error)
	GetByUserName(c context.Context,username string)(*UserResponse,error)
	
}

package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"working.com/bank_dash/internal/domain"
)

// type for working with user repository
type UserRepository struct {
	database   mongo.Database
	collection string
}

// method for creating/inserting user into the database
func (ur *UserRepository) PostUser(c context.Context, userRequest *domain.UserRequest) (*domain.UserResponse, error) {
	collection := ur.database.Collection(ur.collection)
	userId, err := collection.InsertOne(c, userRequest)
	if err != nil {
		return &domain.UserResponse{}, err
	}

	var CreatedUser *domain.UserResponse
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: userId}}).Decode(&CreatedUser)
	if err != nil {
		return &domain.UserResponse{}, err
	}
	return CreatedUser, nil
}

// method for updating the user inside the databse
func (ur *UserRepository) UpdateUser(c context.Context, id string, userRequest *domain.UserRequest) (*domain.UserResponse, error) {
	collection := ur.database.Collection(ur.collection)
	UpdatingUser := bson.M{
		"name":             userRequest.Name,
		"email":            userRequest.Email,
		"dateOfBirth":      userRequest.DateOfBirth,
		"permanentAddress": userRequest.PermanentAddress,
		"postalCode":       userRequest.PostalCode,
		"username":         userRequest.UserName,
		"presentAddress":   userRequest.PresentAddress,
		"city":             userRequest.City,
		"country":          userRequest.Country,
		"profilePicture":   userRequest.ProfilePicture,
	}
	userId, err := collection.UpdateOne(c, bson.D{{Key: "_id", Value: id}}, bson.M{"$set": UpdatingUser})
	if err != nil {
		return &domain.UserResponse{}, err
	}
	var UpdatedUser *domain.UserResponse
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: userId}}).Decode(&UpdatedUser)
	if err != nil {
		return &domain.UserResponse{}, err
	}
	return UpdatedUser, nil

}

// method for updating user preference inside the database
func (ur *UserRepository) UpdatePreference(c context.Context, id string, userPrefernce *domain.UserPreference) (*domain.UserPreference, error) {
	collection := ur.database.Collection(ur.collection)
	UpdatingPreference := bson.M{
		"currency":                     userPrefernce.Currency,
		"sentOrReceiveDigitalCurrency": userPrefernce.SentOrReceiveDigitalCurrency,
		"receiveMerchantOrder":         userPrefernce.ReceiveMerchantOrder,
		"accountRecommendations":       userPrefernce.AccountRecommendations,
		"timeZone":                     userPrefernce.TimeZone,
		"twoFactorAuthentication":      userPrefernce.TwoFactorAuthentication,
	}
	userId, err := collection.UpdateOne(c, bson.D{{Key: "_id", Value: id}}, bson.M{"$set": UpdatingPreference})
	if err != nil {
		return &domain.UserPreference{}, err
	}
	var updatedPreference *domain.UserPreference
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: userId}}).Decode(&updatedPreference)
	if err != nil {
		return &domain.UserPreference{}, err
	}
	return updatedPreference, nil
}

// method for getting user information by using username
func (ur *UserRepository) GetByUserName(c context.Context, username string) (*domain.UserResponse, error) {
	collection := ur.database.Collection(ur.collection)
	var User *domain.UserResponse
	err := collection.FindOne(c, bson.D{{Key: "username", Value: username}}).Decode(&User)
	if err != nil {
		return &domain.UserResponse{}, err
	}
	return User, nil
}

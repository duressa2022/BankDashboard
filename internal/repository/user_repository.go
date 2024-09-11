package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/package/mongo"
)

// type for working with user repository
type UserRepository struct {
	database   mongo.Database
	collection string
}

// method for working with user repository
func NewUserRepository(db mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}

// method for creating/inserting user into the database
func (ur *UserRepository) PostUser(c context.Context, user *domain.User) (*domain.UserResponse, error) {
	collection := ur.database.Collection(ur.collection)
	userId, err := collection.InsertOne(c, user)
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

	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	userResult, err := collection.UpdateOne(c, bson.D{{Key: "_id", Value: Id}}, bson.M{"$set": UpdatingUser})
	if err != nil {
		return nil, errors.New("failed to update")
	}
	if userResult.MatchedCount == 0 {
		return nil, errors.New("no matched document are found")
	}
	if userResult.ModifiedCount == 0 {
		return nil, errors.New("no user information is updated")
	}

	var UpdatedUser *domain.UserResponse
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: Id}}).Decode(&UpdatedUser)
	if err != nil {
		return &domain.UserResponse{}, err
	}
	return UpdatedUser, nil

}

// method for updating user preference inside the database
func (ur *UserRepository) UpdatePreference(c context.Context, id string, userPreference *domain.UserPreference) (*domain.UserResponse, error) {
	collection := ur.database.Collection(ur.collection)

	updatingPreference := bson.M{
		"preference.currency":                     userPreference.Currency,
		"preference.sentOrReceiveDigitalCurrency": userPreference.SentOrReceiveDigitalCurrency,
		"preference.receiveMerchantOrder":         userPreference.ReceiveMerchantOrder,
		"preference.accountRecommendations":       userPreference.AccountRecommendations,
		"preference.timeZone":                     userPreference.TimeZone,
		"preference.twoFactorAuthentication":      userPreference.TwoFactorAuthentication,
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("error while conversion")
	}

	updateResult, err := collection.UpdateOne(c, bson.D{{Key: "_id", Value: objectID}}, bson.M{"$set": updatingPreference})
	if err != nil {
		return nil, errors.New("failed update operation")
	}

	if updateResult.MatchedCount == 0 {
		return nil, errors.New("no matched document found")
	}
	if updateResult.ModifiedCount == 0 {
		return nil, errors.New("no document was modified")
	}

	var updatedUser domain.UserResponse
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: objectID}}).Decode(&updatedUser)
	if err != nil {
		return nil, errors.New("failed to find matched docs")
	}

	return &updatedUser, nil
}

// method for getting user information by using username
func (ur *UserRepository) GetByUserName(c context.Context, username string) (*domain.UserResponse, error) {
	collection := ur.database.Collection(ur.collection)
	var User *domain.UserResponse
	err := collection.FindOne(c, bson.D{{Key: "username", Value: username}}).Decode(&User)
	if err != nil {
		return nil, err
	}
	return User, nil
}

// method for getting user information by using username
func (ur *UserRepository) GetByUserNameForPass(c context.Context, username string) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var User *domain.User
	err := collection.FindOne(c, bson.D{{Key: "username", Value: username}}).Decode(&User)
	if err != nil {
		return nil, err
	}
	return User, nil
}

// method for getting user information by using username
func (ur *UserRepository) GetByUserEmail(c context.Context, email string) (*domain.UserResponse, error) {
	collection := ur.database.Collection(ur.collection)
	var User *domain.UserResponse
	err := collection.FindOne(c, bson.D{{Key: "email", Value: email}}).Decode(&User)
	if err != nil {
		return nil, err
	}
	return User, nil
}

// method for updating password
func (ur *UserRepository) UpdatePassword(c context.Context, username string, passInfo *domain.ChangePassword) error {
	collection := ur.database.Collection(ur.collection)
	var user *domain.User

	err := collection.FindOne(c, bson.D{{Key: "username", Value: username}}).Decode(&user)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passInfo.Password))
	if err != nil {
		return errors.New("password is not correct")
	}

	updated := bson.M{
		"password": passInfo.NewPassword,
	}

	_, err = collection.UpdateOne(c, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$set", Value: updated}})
	return err

}

// method for getting user by using id
func (ur *UserRepository) GetUserId(c context.Context, id string) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	UserId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user *domain.User
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: UserId}}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

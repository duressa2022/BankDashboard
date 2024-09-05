package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"working.com/bank_dash/internal/domain"
)

// type for working with bank service repos
type BankRepository struct {
	database   mongo.Database
	collection string
}

// method for getting banks by using id
func (br *BankRepository) GetBankById(c context.Context, id string) (*domain.BankService, error) {
	collection := br.database.Collection(br.collection)
	company_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.BankService{}, err
	}
	var Bank *domain.BankService
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: company_id}}).Decode(&Bank)
	if err != nil {
		return &domain.BankService{}, err
	}
	return Bank, err
}

// method for updating the bank information based on id
func (br *BankRepository) UpdateBank(c context.Context, id string, bankRequest *domain.BankRequest) (*domain.BankService, error) {
	collection := br.database.Collection(br.collection)
	UpdatingBank := bson.M{
		"name":          bankRequest.Name,
		"details":       bankRequest.Details,
		"numberOfUsers": bankRequest.NumberOfUsers,
		"status":        bankRequest.Status,
		"type":          bankRequest.Type,
		"icon":          bankRequest.Icon,
	}
	BankId, err := collection.UpdateOne(c, bson.D{{Key: "_id", Value: id}}, bson.M{"$set": UpdatingBank})
	if err != nil {
		return nil, err
	}
	var Bank *domain.BankService
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: BankId}}).Decode(&Bank)
	if err != nil {
		return &domain.BankService{}, err
	}
	return Bank, nil
}
// method for deleting bank service from the database
func (br *BankRepository) DeleteBank(c context.Context,id string)error{
	collection:=br.database.Collection(br.collection)
	_,err:=collection.DeleteOne(c,bson.D{{Key: "_id",Value: id}})
	return err
}

// method for posting the bankservice into the database
func (br *BankRepository) PostBank(c context.Context,bank *domain.BankRequest)(*domain.BankService,error){
	collection:=br.database.Collection(br.collection)
	userId,err:=collection.InsertOne(c,bank)
	if err!=nil{
		return nil,err
	}
	var BankService *domain.BankService
	err=collection.FindOne(c,bson.D{{Key: "id",Value: userId}}).Decode(&BankService)
	if err!=nil{
		return nil,err
	}
	return BankService,nil
}

// method for searching by using name of the bank service
func (br *BankRepository) SearchByName(c context.Context,term string)(*domain.BankService,error){
	collection:=br.database.Collection(br.collection)
	var Bank *domain.BankService
	err:=collection.FindOne(c ,bson.D{{Key: "name",Value: term}}).Decode(&Bank)
	if err!=nil{
		return nil,err
	}
	return Bank,nil
}
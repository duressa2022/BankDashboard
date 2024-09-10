package repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/package/mongo"
)

// type for working bank transactions
type TransactionRepo struct {
	database       mongo.Database
	collection     string
	userCollection string
}

// method for working with bank transaction
func NewTransactionRepository(db mongo.Database, transactioncollection string, usercollection string) *TransactionRepo {
	return &TransactionRepo{
		database:       db,
		collection:     transactioncollection,
		userCollection: usercollection,
	}
}

// method for getting transactions based on page and size
func (tr *TransactionRepo) GetTransactions(c context.Context, page int32, size int32) ([]*domain.Transaction, int, error) {
	collection := tr.database.Collection(tr.collection)
	var Transactions []*domain.Transaction
	skip := (page - 1) * size
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(size))
	cursor, err := collection.Find(c, bson.D{{}}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(c)
	for cursor.Next(c) {
		var transaction *domain.Transaction
		err := cursor.Decode(&transaction)
		if err != nil {
			return nil, 0, err
		}
		Transactions = append(Transactions, transaction)
	}

	total, err := collection.CountDocuments(c, bson.D{{}})
	if err != nil {
		return nil, 0, err
	}
	return Transactions, int(total), nil
}

// method for posting transaction information into database
func (tr *TransactionRepo) PostTransaction(c context.Context, Transaction *domain.Transaction) (*domain.Transaction, error) {
	collection := tr.database.Collection(tr.collection)
	transactionId, err := collection.InsertOne(c, Transaction)
	if err != nil {
		return nil, err
	}
	var transaction *domain.Transaction
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: transactionId}}).Decode(&transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil

}

// method for working with deposit
func (tr *TransactionRepo) Deposit(c context.Context, sender string, recipent string, description *domain.DepositTransaction) (*domain.Transaction, error) {
	userCollection := tr.database.Collection(tr.userCollection)

	var senderUser domain.UserResponse
	var recipentUser domain.UserResponse

	err := userCollection.FindOne(c, bson.D{{Key: "username", Value: sender}}).Decode(&sender)
	if err != nil {
		return nil, err
	}
	err = userCollection.FindOne(c, bson.D{{Key: "username", Value: recipent}}).Decode(&recipent)
	if err != nil {
		return nil, err
	}

	if senderUser.AccountBalance < description.Amount {
		return nil, errors.New("amount error")
	}

	senderUpdated := bson.M{
		"accountBalance": senderUser.AccountBalance - description.Amount,
	}
	recipentUpdated := bson.M{
		"accountBalance": recipentUser.AccountBalance + description.Amount,
	}

	_, err = userCollection.UpdateOne(c, bson.D{{Key: "username", Value: sender}}, bson.D{{Key: "$set", Value: senderUpdated}})
	if err != nil {
		return nil, err
	}
	_, err = userCollection.UpdateOne(c, bson.D{{Key: "username", Value: recipent}}, bson.D{{Key: "$set", Value: recipentUpdated}})
	if err != nil {
		return nil, err
	}

	var Transaction domain.Transaction
	Transaction.SenderName = sender
	Transaction.ReceiverUserName = recipent
	Transaction.Type = "transfer"
	Transaction.TransactionId = primitive.NewObjectID()
	Transaction.Description = description.Description
	Transaction.Date = time.Now()

	created, err := tr.PostTransaction(c, &Transaction)
	return created, err

}

// method for getting transaction id
func (tr *TransactionRepo) GetTransactionById(c context.Context, id string) (*domain.Transaction, error) {
	collection := tr.database.Collection(tr.collection)

	var transaction domain.Transaction
	err := collection.FindOne(c, bson.D{{Key: "_id", Value: id}}).Decode(&transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

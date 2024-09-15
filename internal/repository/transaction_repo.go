package repository

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/package/mongo"
)

type TransactionRepository struct {
	database   mongo.Database
	collection string
}

// PostTransaction implements domain.TransactionRepository.
func (t *TransactionRepository) PostTransaction(ctx context.Context, claims jwt.Claims, tr domain.TransactionRequest) (domain.Transaction, error) {
	new_transaction := domain.TransactionDb{
		Type: 					 tr.Type,
		SenderUserName:  claims.(jwt.MapClaims)["username"].(string),
		Description:     tr.Description,
		Date:            time.Now(),
		Amount: 				tr.Amount,
		ReceiverUserName: tr.ReceiverUserName,
	}
	collection := t.database.Collection(t.collection)
	id, err := collection.InsertOne(ctx, new_transaction)

	if err != nil {
		return domain.Transaction{}, err
	}
	var new_tra domain.Transaction
	err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&new_tra)
	if err != nil {
		return domain.Transaction{}, err
	}

	return new_tra, nil

}




// GetTransaction implements domain.TransactionRepository.
func (t *TransactionRepository) GetTransaction(ctx context.Context, claims jwt.Claims, page int, size int) ([]domain.Transaction, int, error) {
	var transactions []domain.Transaction
	collection := t.database.Collection(t.collection)

	skip := (page - 1) * size
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(size))
	findOptions.SetSort(bson.D{{Key: "date", Value: -1}})

	filter := bson.D{{Key: "senderUserName", Value: claims.(jwt.MapClaims)["username"].(string)}}

	totalCount, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return transactions, 0, err
	}
	totalPage := int((totalCount + int64(size) - 1) / int64(size))



	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return transactions, 0, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var transaction domain.Transaction
		if err = cursor.Decode(&transaction); err != nil {
			return transactions, 0, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, totalPage, nil
}

func NewTransactionRepository(database mongo.Database, collection string) domain.TransactionRepository {
	return &TransactionRepository{
		database:   database,
		collection: collection,
	}
}

package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"working.com/bank_dash/internal/domain"
)

// type for working bank transactions
type TransactionRepo struct{
	database mongo.Database
	collection string
}
// method for getting transactions based on page and size
func (tr *TransactionRepo) GetTransactions(c context.Context,page int32,size int32)([]*domain.Transaction,error){
	collection:=tr.database.Collection(tr.collection)
	var Transactions []*domain.Transaction
	skip:=(page-1)*size
	opts:=options.Find().SetSkip(int64(skip)).SetLimit(int64(size))
	cursor,err:=collection.Find(c,bson.D{{}},opts)
	if err!=nil{
		return nil,err
	}
	defer cursor.Close(c)
	for(cursor.Next(c)){
		var transaction *domain.Transaction
		err:=cursor.Decode(&transaction)
		if err!=nil{
			return nil,err
		}
		Transactions=append(Transactions, transaction)
	}
	return Transactions,nil
}
// method for posting transaction information into database
func (tr *TransactionRepo) PostTransaction(c context.Context,TransactionRequest *domain.TransactionRequest)(*domain.Transaction,error){
	collection:=tr.database.Collection(tr.collection)
	transactionId,err:=collection.InsertOne(c,TransactionRequest)
	if err!=nil{
		return nil,err
	}
	var transaction *domain.Transaction
	err=collection.FindOne(c,bson.D{{Key: "_id",Value: transactionId}}).Decode(&transaction)
	if err!=nil{
		return nil,err
	}
	return transaction,nil

}
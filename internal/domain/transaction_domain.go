package domain

import (
	"context"
	"time"
)

// type for working with transaction information
type Transaction struct {
	TransactionId    string    `json:"transactionId" bson:"_transactionId"`
	Type             string    `json:"type" bson:"type"`
	SenderName       string    `json:"senderName" bson:"senderName"`
	Description      string    `json:"description" bson:"description"`
	Date             time.Time `json:"date" bson:"date"`
	Amount           float64   `json:"amount" bson:"amount"`
	ReceiverUserName string    `json:"receiverUserName" bson:"receiverUserName"`
}

// type for working with transaction request body
type TransactionRequest struct {
	Type             string  `json:"type" bson:"type"`
	Description      string  `json:"description" bson:"description"`
	Amount           float64 `json:"amount" bson:"amount"`
	ReceiverUserName string  `json:"receiverUserName" bson:"receiverUserName"`
}

// type for working with transaction request data
type TransactionData struct {
	Time  string `json:"time" bson:"time"`
	Value int64  `json:"value" bson:"value"`
}
// type for working with deposit transaction 
type DepositTransaction struct{
	Description      string  `json:"description" bson:"description"`
	Amount           float64 `json:"amount" bson:"amount"`
}
// interface for working with transaction repo
type TransactionRepo interface{
	GetTransactions(c context.Context,page int64,size int64)([]*Transaction,error)
	PostTransaction(c context.Context,transaction *TransactionRequest)(*Transaction,error)
	Deposit(c context.Context,deposit *DepositTransaction)(*Transaction,error)
	GetTransactionById(c context.Context,id string)(*Transaction,error)
	GetRandomBalanceHistory(c context.Context,time int32)([]*TransactionData,error)
	QuickTransfer(c context.Context,number int32)(*TransactionData,error)
	Incomes(c context.Context,page int32,size int32)([]*Transaction,error)
	Expenses(c context.Context,page int32,size int32)([]*Transaction,error)
	Summary(c context.Context,startDate string ,types string)(*TransactionData,error)
	History(c context.Context)(*TransactionData,error)

}

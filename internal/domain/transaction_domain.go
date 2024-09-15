package domain

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	TransactionCollection = "transaction"
)

type TransactionDb struct {
	Type             string             `json:"type" bson:"type"`
	SenderUserName   string             `json:"senderUserName" bson:"senderUserName"`
	Description      string             `json:"description" bson:"description"`
	Date             time.Time          `json:"date" bson:"date"`
	Amount           float64            `json:"amount" bson:"amount"`
	ReceiverUserName string             `json:"receiverUserName" bson:"receiverUserName"`
}

// type for working with transaction information
type Transaction struct {
	TransactionId    primitive.ObjectID `json:"transactionId" bson:"_id"`
	Type             string             `json:"type" bson:"type"`
	SenderUserName   string             `json:"senderUserName" bson:"senderUserName"`
	Description      string             `json:"description" bson:"description"`
	Date             time.Time          `json:"date" bson:"date"`
	Amount           float64            `json:"amount" bson:"amount"`
	ReceiverUserName string             `json:"receiverUserName" bson:"receiverUserName"`
}

// type for working with transaction request body
type TransactionRequest struct {
	Type             string  `json:"type" bson:"type"`
	Description      string  `json:"description" bson:"description"`
	Amount           float64 `json:"amount" bson:"amount"`
	ReceiverUserName string  `json:"receiverUserName" bson:"receiverUserName"`
}

// type for working with deposit transaction
type TransactionDeposit struct {
	Description string  `json:"description" bson:"description"`
	Amount      float64 `json:"amount" bson:"amount"`
}

// transaction repository interface
type TransactionRepository interface {
	GetTransaction(ctx context.Context, claims jwt.Claims, page int, size int) ([]Transaction, int, error)
	PostTransaction(ctx context.Context, claims jwt.Claims, tr TransactionRequest) (Transaction, error)
	GetTransactionById(ctx context.Context, id primitive.ObjectID) (Transaction, error)
	GetIncomeTransaction(ctx context.Context, claims jwt.Claims, page int, size int) ([]Transaction, int, error)
}

type TransactionUsecase interface {
	GetTransaction(ctx context.Context, claims jwt.Claims, page int, size int) ([]Transaction, int, error)
	PostTransaction(ctx context.Context, claims jwt.Claims, tr TransactionRequest) (Transaction, error)
	DepositTransaction(ctx context.Context, claims jwt.Claims, tr TransactionDeposit) (Transaction, error)
	GetTransactionById(ctx context.Context, id primitive.ObjectID) (Transaction, error)
	GetIncomeTransaction(ctx context.Context, claims jwt.Claims, page int, size int) ([]Transaction, int, error)
}
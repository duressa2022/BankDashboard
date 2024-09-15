package usecase

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"working.com/bank_dash/internal/domain"
)

type TransactionUsecase struct {
	transactionRepository domain.TransactionRepository
	contextTimeout        time.Duration
}

// GetTransactionExpense implements domain.TransactionUsecase.
func (t *TransactionUsecase) GetTransactionExpense(ctx context.Context, claims jwt.Claims, page int, size int) ([]domain.Transaction, int, error) {
	return t.transactionRepository.GetTransactionExpense(ctx, claims, page, size)
}

// GetIncomeTransaction implements domain.TransactionUsecase.
func (t *TransactionUsecase) GetIncomeTransaction(ctx context.Context, claims jwt.Claims, page int, size int) ([]domain.Transaction, int, error) {
	return t.transactionRepository.GetIncomeTransaction(ctx, claims, page, size)
}

// GetTransactionById implements domain.TransactionUsecase.
func (t *TransactionUsecase) GetTransactionById(ctx context.Context, id primitive.ObjectID) (domain.Transaction, error) {
	return t.transactionRepository.GetTransactionById(ctx, id)
}

// DepositTransaction implements domain.TransactionUsecase.
func (t *TransactionUsecase) DepositTransaction(ctx context.Context, claims jwt.Claims, tr domain.TransactionDeposit) (domain.Transaction, error) {
	var trr domain.TransactionRequest
	trr.Amount = tr.Amount
	trr.Description = tr.Description
	trr.ReceiverUserName = claims.(jwt.MapClaims)["username"].(string)
	trr.Type = "deposit"
	return t.transactionRepository.PostTransaction(ctx, claims, trr)
}

// PostTransaction implements domain.TransactionUsecase.
func (t *TransactionUsecase) PostTransaction(ctx context.Context, claims jwt.Claims, tr domain.TransactionRequest) (domain.Transaction, error) {
	return t.transactionRepository.PostTransaction(ctx, claims, tr)
}

// GetTransaction implements domain.TransactionUsecase.
func (t *TransactionUsecase) GetTransaction(ctx context.Context, claims jwt.Claims, page int, size int) ([]domain.Transaction, int, error) {
	return t.transactionRepository.GetTransaction(ctx, claims, page, size)
}

func NewTransactionUsecase(transactionRepository domain.TransactionRepository, contextTimeout time.Duration) domain.TransactionUsecase {
	return &TransactionUsecase{
		transactionRepository: transactionRepository,
		contextTimeout:        contextTimeout,
	}
}

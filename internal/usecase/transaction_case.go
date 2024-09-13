package usecase

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"working.com/bank_dash/internal/domain"
)

type TransactionUsecase struct {
	transactionRepository domain.TransactionRepository
	contextTimeout        time.Duration
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

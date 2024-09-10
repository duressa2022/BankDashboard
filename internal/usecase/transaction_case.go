package usecase

import (
	"context"
	"time"

	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working with the transactions
type TransactionUseCase struct {
	Transaction *repository.TransactionRepo
	Timeout     time.Duration
}

// method for creating new transaction usecase
func NewTransactionUseCase(time time.Duration, transaction *repository.TransactionRepo) *TransactionUseCase {
	return &TransactionUseCase{
		Transaction: transaction,
		Timeout:     time,
	}
}

// method for getting page size and number
func (tu *TransactionUseCase) GetTransactions(c context.Context, page int64, size int64) ([]*domain.Transaction, int, error) {
	return tu.Transaction.GetTransactions(c, int32(page), int32(size))
}

// method for posting the transaction on to the database
func (tu *TransactionUseCase) PostTransaction(c context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	return tu.Transaction.PostTransaction(c, transaction)
}

// method for deposting the transaction on account
func (tu *TransactionUseCase) Deposit(c context.Context, sender string, recipent string, description *domain.DepositTransaction) (*domain.Transaction, error) {
	return tu.Transaction.Deposit(c, sender, recipent, description)
}

// method for getting transaction by using id
func (tu *TransactionUseCase) GetTransactionById(c context.Context, id string) (*domain.Transaction, error) {
	return tu.Transaction.GetTransactionById(c, id)
}

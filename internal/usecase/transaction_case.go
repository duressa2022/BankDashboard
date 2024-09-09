package usecase

import (
	"context"

	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working with the transactions
type TransactionUseCase struct {
	transaction *repository.TransactionRepo
}

// method for creating new transaction usecase
func NewTransactionUseCase(transaction *repository.TransactionRepo) *TransactionUseCase {
	return &TransactionUseCase{
		transaction: transaction,
	}
}

// method for getting page size and number
func (tu *TransactionUseCase) GetTransactions(c context.Context, page int64, size int64) ([]*domain.Transaction, int, error) {
	return tu.transaction.GetTransactions(c, int32(page), int32(size))
}

// method for posting the transaction on to the database
func (tu *TransactionUseCase) PostTransaction(c context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	return tu.transaction.PostTransaction(c, transaction)
}

// method for deposting the transaction on account
func (tu *TransactionUseCase) Deposit(c context.Context, sender string, recipent string, description *domain.DepositTransaction) (*domain.Transaction, error) {
	return tu.transaction.Deposit(c, sender, recipent, description)
}

// method for getting transaction by using id
func (tu *TransactionUseCase) GetTransactionById(c context.Context, id string) (*domain.Transaction, error) {
	return tu.transaction.GetTransactionById(c, id)
}

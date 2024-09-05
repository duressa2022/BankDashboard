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
func (tu *TransactionUseCase) GetTransactions(c context.Context, page int64, size int64) ([]*domain.Transaction, error) {
	return tu.transaction.GetTransactions(c, int32(page), int32(size))
}
// method for posting the transaction on to the database
func (tu *TransactionUseCase)PostTransaction(c context.Context,transaction *domain.TransactionRequest)(*domain.Transaction,error){
	return tu.transaction.PostTransaction(c,transaction)
}

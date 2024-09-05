package usecase

import (
	"context"

	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working with the loan usecase
type LoanUseCase struct {
	loanRepository *repository.LoanRepository
}

// method for creating new loan usecase
func NewLoanUseCase(loan *repository.LoanRepository) *LoanUseCase {
	return &LoanUseCase{
		loanRepository: loan,
	}
}

// method for getting active loans
func (lc *LoanUseCase) ActiveLoan(c context.Context, loan *domain.LoanRequest) (*domain.LoanResponse, error) {
	return lc.loanRepository.ActiveLoan(c, loan)
}

// method for rejecting the loan by id
func (lc *LoanUseCase) Reject(c context.Context, id string) error {
	return lc.loanRepository.Reject(c, id)
}

// method for approving the loan by id
func (lc *LoanUseCase) Approve(c context.Context, id string) (*domain.LoanResponse, error) {
	return lc.loanRepository.Approve(c, id)
}

// method for getting the loan by using id
func (lc *LoanUseCase) GetLoanById(c context.Context, id string) (*domain.LoanResponse, error) {
	return lc.loanRepository.GetLoanById(c, id)
}

// method for getting the loans by using page number and size
func (lc *LoanUseCase) GetMyLoans(c context.Context, page int, size int) ([]*domain.LoanResponse, error) {
	return lc.loanRepository.GetMyLoans(c, page, size)
}

// method for getting the loans by using page number and size
func (lc *LoanUseCase) All(c context.Context, page int, size int) ([]*domain.LoanResponse, error) {
	return lc.loanRepository.GetMyLoans(c, page, size)
}

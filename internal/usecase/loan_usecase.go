package usecase

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working with the loan usecase
type LoanUseCase struct {
	LoanRepository *repository.LoanRepository
	Timeout        time.Duration
}

// method for creating new loan usecase
func NewLoanUseCase(time time.Duration, loan *repository.LoanRepository) *LoanUseCase {
	return &LoanUseCase{
		LoanRepository: loan,
		Timeout:        time,
	}
}

// method for getting active loans
func (lc *LoanUseCase) ActiveLoan(c context.Context, id string, loan *domain.LoanRequest) (*domain.LoanResponse, error) {
	var Loan domain.Loan
	Loan.SerialNumber = primitive.NewObjectID()
	Loan.InterestRate = loan.InterestRate
	Loan.LoanAmount = loan.LoanAmount
	Loan.ActiveLoanStatus = "pending"
	Loan.Duration = loan.Duration
	Loan.AmountLeftToRepay = loan.LoanAmount
	Loan.Type = loan.Type
	Loan.Installment = 0

	UserID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	Loan.UserId = UserID

	return lc.LoanRepository.ActiveLoan(c, &Loan)
}

// method for rejecting the loan by id
func (lc *LoanUseCase) Reject(c context.Context, id string) error {
	return lc.LoanRepository.Reject(c, id)
}

// method for approving the loan by id
func (lc *LoanUseCase) Approve(c context.Context, id string) (*domain.LoanResponse, error) {
	return lc.LoanRepository.Approve(c, id)
}

// method for getting the loan by using id
func (lc *LoanUseCase) GetLoanById(c context.Context, id string) (*domain.LoanResponse, error) {
	return lc.LoanRepository.GetLoanById(c, id)
}

// method for getting the loans by using page number and size
func (lc *LoanUseCase) GetMyLoans(c context.Context, id string, page int, size int) ([]*domain.LoanResponse, int, error) {
	return lc.LoanRepository.GetMyLoans(c, id, page, size)
}

// method for getting the loans by using page number and size
func (lc *LoanUseCase) All(c context.Context, id string, page int, size int) ([]*domain.LoanResponse, int, error) {
	return lc.LoanRepository.GetMyLoans(c, id, page, size)
}

package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	LoanCollection = "loans"
)

// type for working with loan information
type Loan struct {
	LoanAmount        float64            `json:"loanAmount" bson:"loanAmount"`
	AmountLeftToRepay float64            `json:"amountLeftToRepay" bson:"amountLeftToRepay"`
	Duration          int32              `json:"duration" bson:"duration"`
	InterestRate      float64            `json:"interestRate" bson:"interestRate"`
	Installment       int32              `json:"installment" bson:"installment"`
	Type              string             `json:"type" bson:"type"`
	ActiveLoanStatus  string             `json:"activeLoanStatus" bson:"activeLoanStatus"`
	UserId            primitive.ObjectID `json:"userId" bson:"_userId"`
	SerialNumber      primitive.ObjectID `json:"serialNumber" bson:"_serialnumber"`
}

// type for working with loan response
type LoanResponse struct {
	LoanAmount        float64            `json:"loanAmount" bson:"loanAmount"`
	AmountLeftToRepay float64            `json:"amountLeftToRepay" bson:"amountLeftToRepay"`
	Duration          int32              `json:"duration" bson:"duration"`
	InterestRate      float64            `json:"interestRate" bson:"interestRate"`
	Installment       int32              `json:"installment" bson:"installment"`
	Type              string             `json:"type" bson:"type"`
	ActiveLoanStatus  string             `json:"activeLoanStatus" bson:"activeLoanStatus"`
	UserId            primitive.ObjectID `json:"userId" bson:"_userId"`
	SerialNumber      primitive.ObjectID `json:"serialNumber" bson:"_serialnumber"`
}

// type for working with request information
type LoanRequest struct {
	LoanAmount   float64 `json:"loanAmount" bson:"loanAmount"`
	Duration     int32   `json:"duration" bson:"duration"`
	InterestRate float64 `json:"interestRate" bson:"interestRate"`
	Type         string  `json:"type" bson:"type"`
}

// type for working with loan data
type LoanData struct {
	PersonalLoan  int32 `json:"personalLoan" bson:"personalLoan"`
	BusinessLoan  int32 `json:"businessLoan" bson:"businessLoan"`
	CorporateLoan int32 `json:"corporateLoan" bson:"corporateLoan"`
}

// interface for working with Loan repo
type LoanRepo interface {
	ActiveLoan(c context.Context, loan *LoanRequest) (*LoanResponse, error)
	Reject(c context.Context, id string) error
	Approve(c context.Context, id string) (*LoanResponse, error)
	GetLoanById(c context.Context, id string) (*LoanResponse, error)
	GetMyLoans(c context.Context, page int32, size int32) ([]*LoanResponse, error)
	GetDetailLoan(c context.Context) ([]*LoanData, error)
	All(c context.Context, page int32, size int32) ([]*LoanResponse, error)
}

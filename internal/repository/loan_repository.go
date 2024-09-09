package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"working.com/bank_dash/internal/domain"
)

// type for working with loan information
type LoanRepository struct {
	database   mongo.Database
	collection string
}

// method for getting loan based on the information
func (lr *LoanRepository) ActiveLoan(c context.Context, loanRequest *domain.LoanRequest) (*domain.LoanResponse, error) {
	collection := lr.database.Collection(lr.collection)
	filter := bson.M{
		"loanAmount":   loanRequest.LoanAmount,
		"duration":     loanRequest.Duration,
		"interestRate": loanRequest.InterestRate,
		"type":         loanRequest.Type,
	}
	var Loan *domain.LoanResponse
	err := collection.FindOne(c, filter).Decode(&Loan)
	if err != nil {
		return nil, err
	}
	return Loan, nil
}

// method for rejecting the loan requests
func (lr *LoanRepository) Reject(c context.Context, id string) error {
	collection := lr.database.Collection(lr.collection)
	loanId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	updating := bson.M{
		"activeLoanStatus": "rejected",
	}
	_, err = collection.UpdateOne(c, bson.D{{Key: "_id", Value: loanId}}, updating)
	if err != nil {
		return err
	}
	return nil
}

// method for accepting the loan request
func (lr *LoanRepository) Approve(c context.Context, id string) (*domain.LoanResponse, error) {
	collection := lr.database.Collection(lr.collection)
	loanId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	updating := bson.M{
		"activeLoanStatus": "approved",
	}
	_, err = collection.UpdateOne(c, bson.D{{Key: "_id", Value: loanId}}, updating)
	if err != nil {
		return nil, err
	}
	var loan *domain.LoanResponse
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: id}}).Decode(&loan)
	if err != nil {
		return nil, err
	}
	return loan, nil
}

// method for getting loan by using id
func (lr *LoanRepository) GetLoanById(c context.Context, id string) (*domain.LoanResponse, error) {
	collection := lr.database.Collection(lr.collection)
	loanId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var loan *domain.LoanResponse
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: loanId}}).Decode(&loan)
	if err != nil {
		return nil, err
	}
	return loan, err
}

// method for getting by using page and size
func (lr *LoanRepository) GetMyLoans(c context.Context, page int, size int) ([]*domain.LoanResponse, int, error) {
	collection := lr.database.Collection(lr.collection)
	skip := (page - 1) * size
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(size))

	var loans []*domain.LoanResponse
	cursor, err := collection.Find(c, bson.D{{}}, opts)
	if err != nil {
		return nil, 0, err
	}
	for cursor.Next(c) {
		var loan *domain.LoanResponse
		err := cursor.Decode(&loan)
		if err != nil {
			return nil, 0, err
		}
		loans = append(loans, loan)
	}
	total, err := collection.CountDocuments(c, bson.D{{}})
	if err != nil {
		return nil, 0, err
	}
	return loans, int(total), nil
}

// method for getting all loans from the database
func (lr *LoanRepository) All(c context.Context, page int, size int) ([]*domain.LoanResponse, error) {
	collection := lr.database.Collection(lr.collection)
	skip := (page - 1) * size
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(size))

	var loans []*domain.LoanResponse
	cursor, err := collection.Find(c, bson.D{{}}, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(c) {
		var loan *domain.LoanResponse
		err := cursor.Decode(&loan)
		if err != nil {
			return nil, err
		}
		loans = append(loans, loan)
	}
	return loans, nil
}

package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/package/mongo"
)

// type for working with loan information
type LoanRepository struct {
	database   mongo.Database
	collection string
}

// method for working with loan information
func NewLoanRepository(db mongo.Database, collection string) *LoanRepository {
	return &LoanRepository{
		database:   db,
		collection: collection,
	}
}

// method for getting loan based on the information
func (lr *LoanRepository) ActiveLoan(c context.Context, loan *domain.Loan) (*domain.LoanResponse, error) {
	collection := lr.database.Collection(lr.collection)
	loanId, err := collection.InsertOne(c, loan)
	if err != nil {
		return nil, err
	}

	var loanResponse *domain.LoanResponse
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: loanId}}).Decode(&loanResponse)
	if err != nil {
		return nil, err
	}
	return loanResponse, nil
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
	updatedResult, err := collection.UpdateOne(c, bson.D{{Key: "_serialnumber", Value: loanId}}, bson.M{"$set": updating})
	if err != nil {
		return err
	}
	if updatedResult.MatchedCount == 0 {
		return errors.New("no matched document")
	}
	if updatedResult.ModifiedCount == 0 {
		return errors.New("no modified document")
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

	updatedResult, err := collection.UpdateOne(c, bson.D{{Key: "_serialnumber", Value: loanId}}, bson.M{"$set": updating})
	if err != nil {
		return nil, err
	}
	if updatedResult.MatchedCount == 0 {
		return nil, errors.New("no matched docs")
	}
	if updatedResult.ModifiedCount == 0 {
		return nil, errors.New("no modified docs")
	}

	var loan *domain.LoanResponse
	err = collection.FindOne(c, bson.D{{Key: "_serialnumber", Value: loanId}}).Decode(&loan)
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
	err = collection.FindOne(c, bson.D{{Key: "_serialnumber", Value: loanId}}).Decode(&loan)
	if err != nil {
		return nil, err
	}
	return loan, err
}

// method for getting by using page and size
func (lr *LoanRepository) GetMyLoans(c context.Context, id string, page int, size int) ([]*domain.LoanResponse, int, error) {
	collection := lr.database.Collection(lr.collection)
	skip := (page - 1) * size
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(size))

	userid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, 0, err
	}

	var loans []*domain.LoanResponse
	cursor, err := collection.Find(c, bson.D{{Key: "_userId", Value: userid}}, opts)
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

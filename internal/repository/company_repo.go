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

// type for working with company repository
type CompanyRepository struct {
	database   mongo.Database
	collection string
}

// method for company repo
func NewCompanyRepository(db mongo.Database, collection string) *CompanyRepository {
	return &CompanyRepository{
		database:   db,
		collection: collection,
	}
}

// method for getting company by using id
func (cr *CompanyRepository) GetCompanyById(c context.Context, id string) (*domain.Company, error) {
	collection := cr.database.Collection(cr.collection)

	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var company *domain.Company
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: Id}}).Decode(&company)
	if err != nil {
		return nil, err
	}
	return company, err
}

// method for updating the company by using id
func (cr *CompanyRepository) UpdateCompany(c context.Context, id string, company *domain.CompanyRequest) (*domain.Company, error) {
	collection := cr.database.Collection(cr.collection)
	UpdatingCompany := bson.M{
		"companyName": company.CompanyName,
		"type":        company.Type,
		"icon":        company.Icon,
	}

	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	updateResult, err := collection.UpdateOne(c, bson.D{{Key: "_id", Value: Id}}, bson.M{"$set": UpdatingCompany})
	if err != nil {
		return nil, err
	}
	if updateResult.MatchedCount == 0 {
		return nil, errors.New("no matched company is found")

	}
	if updateResult.ModifiedCount == 0 {
		return nil, errors.New("no modified docs")
	}

	var UpdatedCompany *domain.Company
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: Id}}).Decode(&UpdatedCompany)
	if err != nil {
		return nil, err
	}
	return UpdatedCompany, nil
}

// method for deleting company from the database based on id
func (cr *CompanyRepository) DeleteCompany(c context.Context, id string) error {
	collection := cr.database.Collection(cr.collection)

	Id,err:=primitive.ObjectIDFromHex(id)
	if err!=nil{
		return err
	}
	_, err = collection.DeleteOne(c, bson.D{{Key: "_id", Value: Id}})
	return err
}

// method for getting companies based on the page and size
func (cr *CompanyRepository) GetCompanies(c context.Context, page int, size int) ([]*domain.Company, int, error) {
	var Companies []*domain.Company
	collection := cr.database.Collection(cr.collection)

	skip := (page - 1) * size
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(size))
	cursor, err := collection.Find(c, bson.D{{}}, opts)
	if err != nil {
		return nil, 0, err
	}
	for cursor.Next(c) {
		var company *domain.Company
		err := cursor.Decode(&company)
		if err != nil {
			return nil, 0, err
		}
		Companies = append(Companies, company)
	}

	totalnumber, err := collection.CountDocuments(c, bson.D{{}})
	if err != nil {
		return nil, 0, err
	}
	return Companies, int(totalnumber), nil

}

// method for posting company into the database
func (cr *CompanyRepository) PostCompany(c context.Context, company *domain.Company) (*domain.Company, error) {
	collection := cr.database.Collection(cr.collection)

	var companyExisted *domain.Company
	err := collection.FindOne(c, bson.D{{Key: "companyName", Value: company.CompanyName}}).Decode(&companyExisted)
	if err == nil {
		return nil, errors.New("already existing company")
	}

	userId, err := collection.InsertOne(c, company)
	if err != nil {
		return &domain.Company{}, err
	}

	var Company *domain.Company
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: userId}}).Decode(&Company)
	if err != nil {
		return &domain.Company{}, err
	}
	return Company, nil
}

// method for getting a trending companies from the database
func (cr *CompanyRepository) GetTrendingCompanies(c context.Context) ([]*domain.Company, error) {
	collection := cr.database.Collection(cr.collection)
	var companies []*domain.Company
	cursor, err := collection.Find(c, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(c) {
		var company *domain.Company
		err = cursor.Decode(&company)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)

	}
	return companies, nil

}

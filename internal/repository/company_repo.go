package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"working.com/bank_dash/internal/domain"
)

// type for working with company repository
type CompanyRepository struct {
	database   mongo.Database
	collection string
}

// method for getting company by using id
func (cr *CompanyRepository) GetCompanyById(c context.Context, id string) (*domain.Company, error) {
	collection := cr.database.Collection(cr.collection)
	var company *domain.Company
	err := collection.FindOne(c, bson.D{{Key: "_id", Value: id}}).Decode(&company)
	if err != nil {
		return &domain.Company{}, err
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
	_, err := collection.UpdateOne(c, bson.D{{Key: "_id", Value: id}}, bson.M{"$set": UpdatingCompany})
	if err != nil {
		return &domain.Company{}, err
	}
	var UpdatedCompany *domain.Company
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: id}}).Decode(&UpdatedCompany)
	if err != nil {
		return &domain.Company{}, err
	}
	return UpdatedCompany, nil
}

// method for deleting company from the database based on id
func (cr *CompanyRepository) DeleteCompany(c context.Context, id string) error {
	collection := cr.database.Collection(cr.collection)
	_, err := collection.DeleteOne(c, bson.D{{Key: "_id", Value: id}})
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

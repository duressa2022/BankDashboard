package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	CollectionCompany = "company"
)

// type for working with the company domain
type Company struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	CompanyName string             `json:"companyName" bson:"companyName"`
	Type        string             `json:"type" bson:"type"`
	Icon        string             `json:"icon" bson:"icon" `
}

// type for working with companys request
type CompanyRequest struct {
	CompanyName string `json:"companyName" bson:"companyName"`
	Type        string `json:"type" bson:"type"`
	Icon        string `json:"icon" bson:"icon" `
}

// interface for working with companys repo
type CompanyRepository interface {
	GetCompanyById(c context.Context, id string) (*Company, error)
	UpdateCompany(c context.Context, id string, company *CompanyRequest) (*Company, error)
	DeleteCompany(c context.Context, id string) error
	GetCompanies(c context.Context, page int, size int) ([]*Company, error)
	PostCompany(c context.Context, company *CompanyRequest) (*Company, error)
	GetTrendingCompanies(c context.Context) ([]*Company, error)
}

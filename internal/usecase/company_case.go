package usecase

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working with different use cases

type CompanyUseCase struct {
	CompanyRepository *repository.CompanyRepository
	Timeout           time.Duration
}

// method for working with the company repository
func NewcompanyUseCase(time time.Duration, companyRepository *repository.CompanyRepository) *CompanyUseCase {
	return &CompanyUseCase{
		CompanyRepository: companyRepository,
		Timeout:           time,
	}
}

// method for getting company by using id
func (cu *CompanyUseCase) GetCompanyById(c context.Context, id string) (*domain.Company, error) {
	return cu.CompanyRepository.GetCompanyById(c, id)

}

// method for updating company by using id
func (cu *CompanyUseCase) UpdateCompany(c context.Context, id string, company *domain.CompanyRequest) (*domain.Company, error) {
	return cu.CompanyRepository.UpdateCompany(c, id, company)
}

// method deleting the company by using id
func (cu *CompanyUseCase) DeleteCompany(c context.Context, id string) error {
	return cu.CompanyRepository.DeleteCompany(c, id)
}

// method for getting by using page number and size
func (cu *CompanyUseCase) GetCompanies(c context.Context, page int64, size int64) ([]*domain.Company, int, error) {
	return cu.CompanyRepository.GetCompanies(c, int(page), int(size))
}

// method for posting the company on the database
func (cu *CompanyUseCase) PostCompany(c context.Context, company *domain.CompanyRequest) (*domain.Company, error) {
	var newCompany domain.Company
	newCompany.Id = primitive.NewObjectID()
	newCompany.Icon = company.Icon
	newCompany.CompanyName = company.CompanyName
	newCompany.Type = company.Type
	return cu.CompanyRepository.PostCompany(c, &newCompany)
}

// method for getting trending companies
func (cu *CompanyUseCase) GetTrendingCompanies(c context.Context) ([]*domain.Company, error) {
	return cu.CompanyRepository.GetTrendingCompanies(c)
}

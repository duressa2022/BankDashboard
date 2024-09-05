package usecase

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working the company repository
type CompanyUseCase struct {
	companyRepository *repository.CompanyRepository
}

// method for creating new company usecase
func NewCompanyUseCase(company *repository.CompanyRepository) *CompanyUseCase {
	return &CompanyUseCase{
		companyRepository: company,
	}
}

// method for getting company by using id
func (cu *CompanyUseCase) GetCompanyById(c context.Context, id string) (*domain.Company, error) {
	return cu.companyRepository.GetCompanyById(c, id)

}

// method for updating company by using id
func (cu *CompanyUseCase) UpdateCompany(c context.Context, id string, company *domain.CompanyRequest) (*domain.Company, error) {
	return cu.companyRepository.UpdateCompany(c, id, company)
}

// method deleting the company by using id
func (cu *CompanyUseCase) DeleteCompany(c context.Context, id string) error {
	return cu.companyRepository.DeleteCompany(c, id)
}

// method for getting by using page number and size
func (cu *CompanyUseCase) GetCompanies(c context.Context, page int64, size int64) ([]*domain.Company, error) {
	return cu.companyRepository.GetCompanies(c, int(page), int(size))
}

// method for posting the company on the database
func (cu *CompanyUseCase) PostCompany(c context.Context, company *domain.CompanyRequest) (*domain.Company, error) {
	var newCompany domain.Company
	newCompany.Id = primitive.NewObjectID()
	newCompany.Icon = company.Icon
	newCompany.CompanyName = company.CompanyName
	newCompany.Type = company.Type
	return cu.companyRepository.PostCompany(c, &newCompany)
}

// method for getting trending companies 
func (cu *CompanyUseCase) GetTrendingCompanies(c context.Context)([]*domain.Company,error){
	return cu.companyRepository.GetTrendingCompanies(c)
}

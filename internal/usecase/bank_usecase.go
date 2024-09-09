package usecase

import (
	"context"

	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working with bank usecase
type BankUseCase struct {
	bankRepository *repository.BankRepository
}

// method for creating new bank usecase
func NewBankUseCase(bank *repository.BankRepository) *BankUseCase {
	return &BankUseCase{
		bankRepository: bank,
	}
}

// method for getting bank service by using id
func (bu *BankUseCase) GetBankById(c context.Context, id string) (*domain.BankService, error) {
	return bu.bankRepository.GetBankById(c, id)
}

// method for updating bank service by using id
func (bu *BankUseCase) UpdateBank(c context.Context, id string, bank *domain.BankRequest) (*domain.BankService, error) {
	return bu.bankRepository.UpdateBank(c, id, bank)
}

// method for deleting bank service by using id
func (bu *BankUseCase) DeleteBank(c context.Context, id string) error {
	return bu.bankRepository.DeleteBank(c, id)
}

// method for posting bank service on the database
func (bu *BankUseCase) PostBank(c context.Context, bank *domain.BankService) (*domain.BankService, error) {
	return bu.bankRepository.PostBank(c, bank)
}

// methof for bank service by using terms
func (bu *BankUseCase) SearchByName(c context.Context, term string) (*domain.BankService, error) {
	return bu.bankRepository.SearchByName(c, term)
}

// method for get by using page and size number
func (bu *BankUseCase) GetBanks(c context.Context, page int32, size int32) ([]*domain.BankService, int, error) {
	return bu.bankRepository.GetBanks(c, int(page), int(size))
}

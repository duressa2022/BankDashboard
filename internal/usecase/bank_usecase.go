package usecase

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working with bank usecase
type BankUseCase struct {
	BankRepository *repository.BankRepository
	Timeout        time.Duration
}

// method for creating new bank usecase
func NewBankUseCase(time time.Duration, bank *repository.BankRepository) *BankUseCase {
	return &BankUseCase{
		BankRepository: bank,
		Timeout:        time,
	}
}

// method for getting bank service by using id
func (bu *BankUseCase) GetBankById(c context.Context, id string) (*domain.BankService, error) {
	return bu.BankRepository.GetBankById(c, id)
}

// method for updating bank service by using id
func (bu *BankUseCase) UpdateBank(c context.Context, id string, bank *domain.BankRequest) (*domain.BankService, error) {
	return bu.BankRepository.UpdateBank(c, id, bank)
}

// method for deleting bank service by using id
func (bu *BankUseCase) DeleteBank(c context.Context, id string) error {
	return bu.BankRepository.DeleteBank(c, id)
}

// method for posting bank service on the database
func (bu *BankUseCase) PostBank(c context.Context, bank *domain.BankService) (*domain.BankService, error) {
	bank.Id = primitive.NewObjectID()
	return bu.BankRepository.PostBank(c, bank)
}

// methof for bank service by using terms
func (bu *BankUseCase) SearchByName(c context.Context, term string) (*domain.BankService, error) {
	return bu.BankRepository.SearchByName(c, term)
}

// method for get by using page and size number
func (bu *BankUseCase) GetBanks(c context.Context, page int32, size int32) ([]*domain.BankService, int, error) {
	return bu.BankRepository.GetBanks(c, int(page), int(size))
}

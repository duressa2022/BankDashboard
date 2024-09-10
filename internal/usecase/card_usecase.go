package usecase

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working with card usecase
type CardUseCase struct {
	CardRepository *repository.CardRepository
	Timeout        time.Duration
}

// method for creating new card usecase
func NewCardUseCase(time time.Duration, card *repository.CardRepository) *CardUseCase {
	return &CardUseCase{
		CardRepository: card,
		Timeout:        time,
	}
}

// method for getting by using page number and size
func (cu *CardUseCase) GetCards(c context.Context, page int32, size int32) ([]*domain.CardResponse, int, error) {
	return cu.CardRepository.GetCards(c, page, size)
}

// method for posting cards/information on the database
func (cu *CardUseCase) PostCard(c context.Context, id string ,card *domain.CardRequest) (*domain.Card, error) {
	var Card domain.Card
	Card.Balance = card.Balance
	Card.CardHolder = card.CardHolder
	Card.CardNumber = "car number"
	Card.CardType = card.CardType
	Card.Passcode = card.Passcode
	Card.ExpiryDate = card.ExpiryDate
	Card.Id = primitive.NewObjectID()

	UserId,err:=primitive.ObjectIDFromHex(id)
	if err!=nil{
		return nil,err
	}
	Card.UserId=UserId
	return cu.CardRepository.PostCard(c, &Card)
}

// method for getting card by using id
func (cu *CardUseCase) GetCardById(c context.Context, id string) (*domain.Card, error) {
	return cu.CardRepository.GetCardById(c, id)
}

// method for deleting card by using id
func (cu *CardUseCase) DeleteCard(c context.Context, id string) error {
	return cu.CardRepository.DeleteCard(c, id)
}

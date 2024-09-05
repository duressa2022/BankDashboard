package usecase

import (
	"context"

	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working with card usecase
type CardUseCase struct{
	cardRepository *repository.CardRepository
}

// method for creating new card usecase
func NewCardUseCase(card *repository.CardRepository)*CardUseCase{
	return &CardUseCase{
		cardRepository: card,
	}
}
// method for getting by using page number and size
func (cu *CardUseCase)GetCards(c context.Context,page int32,size int32)([]*domain.CardResponse,error){
	return cu.cardRepository.GetCards(c,page,size)
}
// method for posting cards/information on the database 
func (cu *CardUseCase)PostCard(c context.Context,card *domain.CardRequest)(*domain.Card,error){
	return cu.cardRepository.PostCard(c,card)
}
// method for getting card by using id
func (cu *CardUseCase)GetCardById(c context.Context,id string)(*domain.Card,error){
	return cu.cardRepository.GetCardById(c,id)
}
// method for deleting card by using id
func (cu *CardUseCase)DeleteCard(c context.Context,id string)error{
	return cu.cardRepository.DeleteCard(c,id)
}

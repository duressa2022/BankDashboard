package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CardCollection = "cards"
)
// type for working with card information
type Card struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	Balance    float64            `json:"balance" bson:"balance"`
	CardHolder string             `json:"cardHolder" bson:"cardHolder"`
	ExpiryDate time.Time          `json:"expiryDate" bson:"expiryDate"`
	CardNumber string             `json:"cardNumber" bson:"cardNumber"`
	Passcode   string             `json:"passcode" bson:"passcode"`
	CardType   string             `json:"cardType" bson:"cardType"`
	UserId     primitive.ObjectID `json:"userId" bson:"_userId"`
}

// type for working with card request
type CardRequest struct {
	Balance    float64   `json:"balance" bson:"balance"`
	CardHolder string    `json:"cardHolder" bson:"cardHolder"`
	ExpiryDate time.Time `json:"expiryDate" bson:"expiryDate"`
	Passcode   string    `json:"passcode" bson:"passcode"`
	CardType   string    `json:"cardType" bson:"cardType"`
}

// type for working with card response
type CardResponse struct {
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	Balance        float64            `json:"balance" bson:"balance"`
	CardHolder     string             `json:"cardHolder" bson:"cardHolder"`
	ExpiryDate     time.Time          `json:"expiryDate" bson:"expiryDate"`
	SemiCardNumber string             `json:"semiCardNumber" bson:"cardNumber"`
	CardType       string             `json:"cardType" bson:"cardType"`
}

// interface for working card repo

type CardRepo interface {
	GetCards(c context.Context, page int32, size int32) ([]*CardResponse, error)
	PostCard(c context.Context, card *CardRequest) (*Card, error)
	GetCardById(c context.Context, id string) (*Card, error)
	DeleteCard(c context.Context, id string) error
}

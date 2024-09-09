package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"working.com/bank_dash/internal/domain"
)

// type for working with card information
type CardRepository struct {
	database  mongo.Database
	collecton string
}

// method for getting cards based on page and size
func (cr *CardRepository) GetCards(c context.Context, page int32, size int32) ([]*domain.CardResponse, int, error) {
	collection := cr.database.Collection(cr.collecton)
	var Cards []*domain.CardResponse
	skip := (page - 1) * size
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(size))

	curser, err := collection.Find(c, bson.D{{}}, opts)
	if err != nil {
		return nil, 0, err
	}
	for curser.Next(c) {
		var card *domain.CardResponse
		err := curser.Decode(&card)
		if err != nil {
			return nil, 0, err
		}
		Cards = append(Cards, card)
	}

	total, err := collection.CountDocuments(c, bson.D{{}})
	if err != nil {
		return nil, 0, err
	}
	return Cards, int(total), nil

}

// method for posting card on the database
func (cr *CardRepository) PostCard(c context.Context, cardRequest *domain.Card) (*domain.Card, error) {
	collection := cr.database.Collection(cr.collecton)
	cardId, err := collection.InsertOne(c, cardRequest)
	if err != nil {
		return nil, err
	}
	var Card *domain.Card
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: cardId}}).Decode(&Card)
	if err != nil {
		return nil, err
	}
	return Card, nil
}

// method for getting card based on the card id
func (cr *CardRepository) GetCardById(c context.Context, id string) (*domain.Card, error) {
	collection := cr.database.Collection(cr.collecton)
	card_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var card *domain.Card
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: card_id}}).Decode(&card)
	if err != nil {
		return nil, err
	}
	return card, nil
}

// method for deleting card by using card id
func (cr *CardRepository) DeleteCard(c context.Context, id string) error {
	collection := cr.database.Collection(cr.collecton)
	card_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.D{{Key: "_id", Value: card_id}})
	return err
}

package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/package/mongo"
)

// type for working on the chat repository
type ChatRepository struct {
	database   mongo.Database
	collection string
}

// method for creating new chat repository
func NewChatRepository(db mongo.Database, collection string) *ChatRepository {
	return &ChatRepository{
		database:   db,
		collection: collection,
	}
}

// method for storing chat message into the database
func (cr *ChatRepository) StoreMessage(c context.Context, message *domain.ChatMessage) (*domain.ChatMessage, error) {
	collection := cr.database.Collection(cr.collection)
	chatId, err := collection.InsertOne(c, message)
	if err != nil {
		return nil, err
	}

	var chatData domain.ChatMessage
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: chatId}}).Decode(&chatData)
	if err != nil {
		return nil, err
	}

	return &chatData, nil
}

// method getting chat history from the database based id
func (cr *ChatRepository) GetMessage(c context.Context, id string) ([]*domain.ChatMessage, error) {
	collection := cr.database.Collection(cr.collection)
	UserId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(c, bson.D{{Key: "_userid", Value: UserId}})
	if err != nil {
		return nil, err
	}

	var chatHistroy []*domain.ChatMessage
	for cursor.Next(c) {
		var histroy *domain.ChatMessage
		err := cursor.Decode(&histroy)
		if err != nil {
			return nil, err
		}
		chatHistroy = append(chatHistroy, histroy)
	}
	return chatHistroy, nil
}

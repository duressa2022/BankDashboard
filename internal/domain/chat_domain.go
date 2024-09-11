package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type for working with chatting domain
type ChatMessage struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"userid" bson:"_userId"`
	Message   string             `json:"message" bson:"message"`
	Response  string             `json:"response" bson:"response"`
	TimeStamp time.Time          `json:"timeStamp" bson:"timeStamp"`
}

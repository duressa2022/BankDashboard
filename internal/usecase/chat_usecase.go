package usecase

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
)

// type for working chat usecase
type ChatUseCase struct {
	ChatRepository *repository.ChatRepository
	Timeout        time.Duration
}

// method for creating new chat usecase
func NewChatUseCase(time time.Duration, chatRepository *repository.ChatRepository) *ChatUseCase {
	return &ChatUseCase{
		ChatRepository: chatRepository,
		Timeout:        time,
	}
}

// method for storing chat in the database
func (cu *ChatUseCase) StoreMessage(c context.Context, id string, message *domain.ChatMessage) (*domain.ChatResponse, error) {
	message.ID = primitive.NewObjectID()

	UserId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	message.UserID = UserId
	message.TimeStamp = time.Now()
	return cu.ChatRepository.StoreMessage(c, message)
}

// method for creating prompt based on the history
func (cu *ChatUseCase) CreatePrompt(c context.Context, id string, message *domain.ChatRequest) (string, error) {
	prompt := "This the conversion between the ai assistance and the user.\n"

	chatHistory, err := cu.ChatRepository.GetMessage(c, id)
	if err != nil {
		return "", err
	}

	for _, chat := range chatHistory {
		prompt += fmt.Sprintf("user: %s\n assistant: %s\n", chat.Message, chat.Response)
	}
	prompt += fmt.Sprintf("User: %s\n", message.Message)
	return prompt, nil
}

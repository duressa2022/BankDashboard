package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"working.com/bank_dash/config"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/usecase"
)

// type for working chat controller
type ChatController struct {
	ChatUseCase *usecase.ChatUseCase
	Env         *config.Env
}

// method for working with chat controller
func NewChatController(env *config.Env, chatusecase *usecase.ChatUseCase) *ChatController {
	return &ChatController{
		ChatUseCase: chatusecase,
		Env:         env,
	}
}

// handler for working with request and response with the bot
func (cc *ChatController) HandleChat(c *gin.Context) {
	var Message *domain.ChatRequest
	if err := c.BindJSON(&Message); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Invalid data for json"})
		return
	}

	Userid, exist := c.Get("id")
	if !exist {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data for userid"})
		return
	}

	userID, okay := Userid.(string)
	if !okay {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Invalid data type for userid"})
		return
	}

	prompt, err := cc.ChatUseCase.CreatePrompt(c, userID, Message)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	client := resty.New()
	response, err := client.R().
		SetHeader("content-type", "application/json").
		SetHeader("Authorization", "Bearer "+cc.Env.API).
		SetBody(map[string]interface{}{
			"message": prompt,
		}).
		Post("https://api.gemini.com/chat")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error while messaging with the bot"})
		return
	}

	var chatMessage domain.ChatMessage
	chatMessage.Message = Message.Message
	chatMessage.Response = response.String()
	cc.ChatUseCase.StoreMessage(c, userID, &chatMessage)

	BotResponse := map[string]interface{}{
		"success":  true,
		"message":  "message from the assistant",
		"response": response.String(),
	}
	c.IndentedJSON(http.StatusOK, BotResponse)

}

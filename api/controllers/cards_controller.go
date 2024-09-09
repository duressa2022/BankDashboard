package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/config"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/usecase"
)

// struct for working with card controller
type CardController struct {
	cardUseCase *usecase.CardUseCase
	Env         *config.Env
}

// method for working card controller
func NewCardController(env *config.Env, card *usecase.CardUseCase) *CardController {
	return &CardController{
		cardUseCase: card,
		Env:         env,
	}
}

// handler for working with card pagination
func (cc *CardController) GetCards(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")

	pageNumber, err := strconv.Atoi(page)
	if err != nil || pageNumber < 1 {
		pageNumber = 1
	}
	sizeNumber, err := strconv.Atoi(size)
	if err != nil || sizeNumber < 1 {
		sizeNumber = 1
	}

	cards, total, err := cc.cardUseCase.GetCards(c, int32(pageNumber), int32(sizeNumber))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error "})
		return
	}

	response := map[string]interface{}{
		"succces": true,
		"message": "cards",
		"content": map[string]interface{}{
			"data": cards,
		},
		"totalpages": total,
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for posting card
func (cc *CardController) PostCard(c *gin.Context) {
	var card domain.Card
	if err := c.BindJSON(&card); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid data"})
		return
	}

	created, err := cc.cardUseCase.PostCard(c, &card)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "created",
		"data":    created,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler for getting card by id
func (cc *CardController) GetCardById(c *gin.Context) {
	id := c.Query("id")

	card, err := cc.cardUseCase.GetCardById(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "card",
		"data":    card,
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for getting card by id
func (cc *CardController) GetDeleteById(c *gin.Context) {
	id := c.Query("id")

	err := cc.cardUseCase.DeleteCard(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "card",
		"data":    map[string]interface{}{},
	}
	c.IndentedJSON(http.StatusOK, response)

}

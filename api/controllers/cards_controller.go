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
	CardUseCase *usecase.CardUseCase
	Env         *config.Env
}

// method for working card controller
func NewCardController(env *config.Env, card *usecase.CardUseCase) *CardController {
	return &CardController{
		CardUseCase: card,
		Env:         env,
	}
}

// handler for working with card pagination
func (cc *CardController) GetCards(c *gin.Context) {
	useId,exist:=c.Get("id")
	if !exist{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"error for data"})
		return 
	}
	Id,okay:=useId.(string)
	if !okay{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"error of type"})
		return 
	}

	page := c.Query("page")
	size := c.Query("size")

	pageNumber, err := strconv.Atoi(page)
	if err != nil || pageNumber < 1 {
		pageNumber = 1
	}
	sizeNumber, err := strconv.Atoi(size)
	if err != nil || sizeNumber < 1 {
		sizeNumber = 10
	}

	cards, total, err := cc.CardUseCase.GetCards(c, Id,int32(pageNumber), int32(sizeNumber))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error "})
		return
	}
	totalPages := (total + sizeNumber - 1) / sizeNumber

	response := map[string]interface{}{
		"succces": true,
		"message": "cards",
		"content": map[string]interface{}{
			"data": cards,
		},
		"totalpages": totalPages,
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for posting card
func (cc *CardController) PostCard(c *gin.Context) {
	var card domain.CardRequest
	if err := c.BindJSON(&card); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid data"})
		return
	}

	UserId, exist := c.Get("id")
	if !exist {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid data"})
		return
	}

	Id, okay := UserId.(string)
	if !okay {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	created, err := cc.CardUseCase.PostCard(c, Id, &card)
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
	id := c.Param("id")

	card, err := cc.CardUseCase.GetCardById(c, id)
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
func (cc *CardController) DeleteById(c *gin.Context) {
	id := c.Param("id")

	err := cc.CardUseCase.DeleteCard(c, id)
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

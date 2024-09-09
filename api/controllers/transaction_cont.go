package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/config"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/usecase"
)

// struct for working with transaction usecases
type TransactionController struct {
	TransactionUseCase *usecase.TransactionUseCase
	Env                *config.Env
}

// method for working transaction controller
func NewTransactionController(env *config.Env, transaction *usecase.TransactionUseCase) *TransactionController {
	return &TransactionController{
		TransactionUseCase: transaction,
		Env:                env,
	}
}

// handler for working with paginations
func (tc *TransactionController) GetTransactionsByLimit(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")

	pageNUmber, err := strconv.Atoi(page)
	if err != nil || pageNUmber < 1 {
		pageNUmber = 1
	}
	sizeNumber, err := strconv.Atoi(size)
	if err != nil || sizeNumber < 1 {
		sizeNumber = 10
	}

	transaction, totalpages, err := tc.TransactionUseCase.GetTransactions(c, int64(pageNUmber), int64(sizeNumber))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "transactions",
		"content": map[string]interface{}{
			"data": transaction,
		},
		"totalpages": totalpages,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler for posting transaction
func (bc *TransactionController) PostTransaction(c *gin.Context) {
	var transaction domain.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	created, err := bc.TransactionUseCase.PostTransaction(c, &transaction)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "posted",
		"data":    created,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler conducting the transaction
func (bc *TransactionController) Deposit(c *gin.Context) {
	senderUser, okay := c.Get("username")
	if !okay {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	sender, err := senderUser.(string)
	if !err {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of data"})
		return
	}
	recipent := c.Query("username")

	var description domain.DepositTransaction
	if err := c.BindJSON(&description); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}

	transaction, errr := bc.TransactionUseCase.Deposit(c, sender, recipent, &description)
	if errr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": errr.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "completed",
		"data":    transaction,
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for getting transaction by id
func (bc *TransactionController) GetTransactionById(c *gin.Context) {
	id := c.Query("id")

	transaction, err := bc.TransactionUseCase.GetTransactionById(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "transaction",
		"data":    transaction,
	}
	c.IndentedJSON(http.StatusOK, response)

}

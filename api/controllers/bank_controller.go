package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/config"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/usecase"
)

// struct for working with bank controller
type BankController struct {
	BankUseCase *usecase.BankUseCase
	Env         *config.Env
}

// function for working with bank controller
func NewBankController(env *config.Env, bank *usecase.BankUseCase) *BankController {
	return &BankController{
		BankUseCase: bank,
		Env:         env,
	}
}

// handler for getting service by id
func (bc *BankController) GetBankByID(c *gin.Context) {
	Id := c.Param("id")

	bank, err := bc.BankUseCase.GetBankById(c, Id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "your bank is:",
		"data":    bank,
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for updating bank information
func (bc *BankController) UpdateBank(c *gin.Context) {
	var bank domain.BankRequest
	if err := c.BindJSON(&bank); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	id := c.Param("id")

	updated, err := bc.BankUseCase.UpdateBank(c, id, &bank)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	response := map[string]interface{}{
		"success": true,
		"message": "updated",
		"data":    updated,
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for deleting bank service
func (bc *BankController) DeleteBank(c *gin.Context) {
	id := c.Param("id")

	err := bc.BankUseCase.DeleteBank(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "deleted",
		"data":    map[string]interface{}{},
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler for working with pagination
func (bc *BankController) GetBankByLimit(c *gin.Context) {
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

	banks, total, err := bc.BankUseCase.GetBanks(c, int32(pageNumber), int32(sizeNumber))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	totalPages := (total + sizeNumber - 1) / sizeNumber
	response := map[string]interface{}{
		"success": true,
		"message": "banks",
		"content": map[string]interface{}{
			"data": banks,
		},
		"totalpages": totalPages,
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for working with posting the bank
func (bc *BankController) PostBank(c *gin.Context) {
	var bank *domain.BankService
	if err := c.BindJSON(&bank); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	created, err := bc.BankUseCase.PostBank(c, bank)
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

//handler for searching bank service by using name
func (bc *BankController) SearchByName(c *gin.Context) {
	searchTerm := c.Query("query")

	result, err := bc.BankUseCase.SearchByName(c, searchTerm)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "search result",
		"data":    result,
	}
	c.IndentedJSON(http.StatusOK, response)
}

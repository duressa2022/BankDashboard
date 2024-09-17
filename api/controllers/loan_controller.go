package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/config"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/usecase"
)

// struct for working with loan controller
type LoanController struct {
	LoanUseCase *usecase.LoanUseCase
	Env         *config.Env
}

// method for working with loan controller
func NewLoanController(env *config.Env, loan *usecase.LoanUseCase) *LoanController {
	return &LoanController{
		LoanUseCase: loan,
		Env:         env,
	}
}

// handler for working posting loan
func (lc *LoanController) ActiveLoan(c *gin.Context) {
	var loan domain.LoanRequest
	if err := c.BindJSON(&loan); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Invalid data"})
		return
	}

	Id, exist := c.Get("id")
	if !exist {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error od data"})
		return
	}

	UserID, okay := Id.(string)
	if !okay {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of data type"})
		return
	}

	Active, err := lc.LoanUseCase.ActiveLoan(c, UserID, &loan)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "active loans",
		"data":    Active,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler for rejecting loans
func (lc *LoanController) Reject(c *gin.Context) {
	id := c.Param("id")

	err := lc.LoanUseCase.Reject(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "rejected",
		"data":    map[string]interface{}{},
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for approving loans
func (lc *LoanController) Approve(c *gin.Context) {
	id := c.Param("id")

	loan, err := lc.LoanUseCase.Approve(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "approved",
		"data":    loan,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler for getting loan by using id
func (lc *LoanController) GetLoanById(c *gin.Context) {
	id := c.Param("id")

	loan, err := lc.LoanUseCase.GetLoanById(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "loan",
		"data":    loan,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler for working with loan pagination
func (lc *LoanController) GetMyLoans(c *gin.Context) {
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

	id, exist := c.Get("id")
	if !exist {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "message": "error of data"})
		return
	}
	UserId, okay := id.(string)
	if !okay {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "message": "error of type"})
		return
	}

	loans, total, err := lc.LoanUseCase.GetMyLoans(c, UserId, pageNumber, sizeNumber)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	totalpages := (total + sizeNumber - 1) / sizeNumber

	response := map[string]interface{}{
		"success": true,
		"message": "loans",
		"content": map[string]interface{}{
			"data": loans,
		},
		"totalPages": totalpages,
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for working with loan pagination
func (lc *LoanController) All(c *gin.Context) {
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

	id, exist := c.Get("id")
	if !exist {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "message": "error of data"})
		return
	}

	userid, okay := id.(string)
	if !okay {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "message": "error of type"})
		return
	}

	loans, total, err := lc.LoanUseCase.GetMyLoans(c, userid, pageNumber, sizeNumber)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	totalpages := (total + sizeNumber - 1) / sizeNumber

	response := map[string]interface{}{
		"success": true,
		"message": "loans",
		"content": map[string]interface{}{
			"data": loans,
		},
		"totalPages": totalpages,
	}
	c.IndentedJSON(http.StatusOK, response)

}

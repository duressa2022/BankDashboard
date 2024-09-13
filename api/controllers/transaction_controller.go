package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/config"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/package/tokens"
)

type TransactionController struct {
	TransactionUsecase domain.TransactionUsecase
	Env *config.Env
}

func (tc *TransactionController) GetTransaction(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(token, "Bearer ")
	p := c.Query("page")
	s := c.Query("size")
	page, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "page must be a number",
		})
		return
	}
	size, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "size must be a number",
		})
		return
	}

	if page < 1 || size < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "page and size must be greater than 0",
		})
		return
	}

	
	claims, _ := tokens.GetUserClaims(tokenString, tc.Env.AccessTokenSecret)
	data, totalPage, err := tc.TransactionUsecase.GetTransaction(c, claims, int(page), int(size));
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("data: ", data)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "transaction fetched successfully",
		"data": gin.H{
			"content": data,
		},
		"totalPage": totalPage,
	})
}

func (tc *TransactionController) PostTransaction(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(token, "Bearer ")
	var tr domain.TransactionRequest
	err := c.BindJSON(&tr)
	fmt.Println("PostTransaction tr: ", tr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	claims, _ := tokens.GetUserClaims(tokenString, tc.Env.AccessTokenSecret)
	data, err := tc.TransactionUsecase.PostTransaction(c, claims, tr);
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "transaction created successfully",
		"data": gin.H{
			"content": data,
		},
	})
}
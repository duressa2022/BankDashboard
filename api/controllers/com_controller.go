package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/config"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/usecase"
)

type CompanyController struct {
	CompanyUseCase *usecase.CompanyUseCase
	Env            *config.Env
}

func NewCompanyrController(env *config.Env, companyCase *usecase.CompanyUseCase) *CompanyController {
	return &CompanyController{
		CompanyUseCase: companyCase,
		Env:            env,
	}
}

// handler for getting company information by id
func (cc *CompanyController) GetCompanyByID(c *gin.Context) {
	Id := c.Param("id")

	company, err := cc.CompanyUseCase.GetCompanyById(c, Id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	response := map[string]interface{}{
		"success": true,
		"message": "company you needed",
		"data":    company,
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for updating the company by using id
func (cc *CompanyController) UpdateCompanyByID(c *gin.Context) {
	var company domain.CompanyRequest
	if err := c.BindJSON(&company); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	Id := c.Param("id")

	updated, err := cc.CompanyUseCase.UpdateCompany(c, Id, &company)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	response := map[string]interface{}{
		"succes":  true,
		"message": updated,
		"data":    updated,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler for deleting company by using id
func (cc *CompanyController) DeleteCompanyByID(c *gin.Context) {
	Id := c.Param("id")

	err := cc.CompanyUseCase.DeleteCompany(c, Id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "company is deleted",
		"data":    map[string]interface{}{},
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for working with pagination
func (cc *CompanyController) GetCompaniessBYLimit(c *gin.Context) {
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

	companies, total, err := cc.CompanyUseCase.GetCompanies(c, int64(pageNumber), int64(sizeNumber))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	totalpages := (total + sizeNumber - 1) / sizeNumber
	response := map[string]interface{}{
		"success": true,
		"message": "companies",
		"data": map[string]interface{}{
			"content":    companies,
			"totalpages": totalpages,
		},
	}
	c.IndentedJSON(http.StatusOK, response)

}

// handler for posting company
func (cc *CompanyController) GetAllCompany(c *gin.Context) {
	var company domain.CompanyRequest
	if err := c.BindJSON(&company); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error of data"})
		return
	}
	companyInfo, err := cc.CompanyUseCase.PostCompany(c, &company)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	response := map[string]interface{}{
		"success": true,
		"message": "company is created",
		"data":    companyInfo,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// handler for getting trending comapnies
func (cc *CompanyController) GetTrendingCompanies(c *gin.Context) {
	companies, err := cc.CompanyUseCase.GetTrendingCompanies(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	response := map[string]interface{}{
		"success": true,
		"message": "trending companies",
		"data":    companies,
	}
	c.IndentedJSON(http.StatusOK, response)
}

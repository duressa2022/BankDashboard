package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/api/controllers"
	"working.com/bank_dash/config"
	"working.com/bank_dash/internal/domain"
	"working.com/bank_dash/internal/repository"
	"working.com/bank_dash/internal/usecase"
	"working.com/bank_dash/package/mongo"
)

// method for init protected routes for company
func initProtectedCompanyRoute(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCompanyRepository(db, domain.CollectionCompany)
	cc := controllers.CompanyController{
		CompanyUseCase: usecase.NewcompanyUseCase(timeout, ur),
		Env:            env,
	}
	group.GET("/:id", cc.GetCompanyByID)
	group.PUT("/:id", cc.UpdateCompanyByID)
	group.DELETE("/:id", cc.DeleteCompanyByID)
	group.GET("", cc.GetCompaniessBYLimit)
	group.POST("/", cc.PostCompany)
	group.GET("/trending-companies", cc.GetTrendingCompanies)
}

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


// method for init protected route for bank service
func initProtectedBankRoute(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	br := repository.NewBankRepository(db, domain.CollectionBank)
	cr := controllers.BankController{
		BankUseCase: usecase.NewBankUseCase(timeout, br),
		Env:         env,
	}
	group.GET("/:id", cr.GetBankByID)
	group.PUT("/:id", cr.UpdateBank)
	group.DELETE("/:id", cr.DeleteBank)
	group.GET("/", cr.GetBankByLimit)
	group.POST("/", cr.PostBank)
	group.GET("/search", cr.SearchByName)

}

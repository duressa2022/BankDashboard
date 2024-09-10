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
// method for init protected route for transaction
func initProtectedTransactionRoute(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	tr := repository.NewTransactionRepository(db, domain.TransactionCollection, domain.CollectionUser)
	tc := controllers.TransactionController{
		TransactionUseCase: usecase.NewTransactionUseCase(timeout, tr),
		Env:                env,
	}
	group.GET("", tc.GetTransactionsByLimit)
	group.POST("", tc.PostTransaction)
	group.POST("/deposit", tc.Deposit)
	group.GET("/:id", tc.GetTransactionById)
}

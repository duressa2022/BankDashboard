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
	// "go.mongodb.org/mongo-driver/mongo"
)

// method for init protected route for transaction
func initProtectedTransactionRoute(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	
	tr := repository.NewTransactionRepository(db, domain.TransactionCollection)

	tc := controllers.TransactionController{
		TransactionUsecase: usecase.NewTransactionUsecase(tr, timeout),
		Env: env,
	}

	group.GET("", tc.GetTransaction)
	group.POST("", tc.PostTransaction)
	group.POST("/deposit", tc.DepositTransaction)
	// group.GET("/:id", tc.GetTransactionById)
}

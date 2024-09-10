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

// method for init protected route for loan
func initProtectedLoanRoute(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	lr := repository.NewLoanRepository(db, domain.LoanCollection)
	lc := controllers.LoanController{
		LoanUseCase: usecase.NewLoanUseCase(timeout, lr),
		Env:         env,
	}
	group.POST("", lc.ActiveLoan)
	group.POST("/:id/reject", lc.Reject)
	group.POST("/:id/approve", lc.Approve)
	group.GET("/:id", lc.GetLoanById)
	group.GET("/my-loans", lc.GetMyLoans)
	group.GET("/loans", lc.All)
}

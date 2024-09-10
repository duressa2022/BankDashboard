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

// method for init protected route for cards
func initProtectedCardRoute(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	cr := repository.NewCardRepository(db, domain.CardCollection)
	cc := controllers.CardController{
		CardUseCase: usecase.NewCardUseCase(timeout, cr),
		Env:         env,
	}
	group.GET("/", cc.GetCards)
	group.POST("/", cc.PostCard)
	group.GET("/:id", cc.GetCardById)
	group.DELETE("/:id", cc.DeleteById)
}


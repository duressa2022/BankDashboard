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

// method for init public route for the chat controller
func initPublicChatRoute(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	cr := repository.NewChatRepository(db, domain.ChatCollection)
	cc := controllers.ChatController{
		ChatUseCase: usecase.NewChatUseCase(timeout, cr),
		Env:         env,
	}
	group.POST("/chat", cc.HandleChat)
}

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

// method for init public route for user
func initPublicUserRoutes(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	uc := controllers.UserController{
		UserUseCase: usecase.NewUserUseCase(timeout, ur),
		Env:         env,
	}
	group.POST("/register", uc.Signup)
	group.POST("/login", uc.LoginIN)
	group.POST("/refresh_token", uc.RefreshToken)
	group.POST("/change_password", uc.ChangePassword)
}

// method for init protected for user
func initProtectedUserRoutes(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	uc := controllers.UserController{
		UserUseCase: usecase.NewUserUseCase(timeout, ur),
		Env:         env,
	}
	group.PUT("/update", uc.UpdateProfile)
	group.PUT("/update-preference", uc.UpdatePreference)
	group.GET("/:username", uc.GetByUserName)
	group.GET("/current", uc.GetCurrentUser)
}

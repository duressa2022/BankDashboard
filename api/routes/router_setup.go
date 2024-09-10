package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/api/middlewares"
	"working.com/bank_dash/config"
	"working.com/bank_dash/package/mongo"
)


// method for setting the route
func SetUpRoute(env *config.Env, timeout time.Duration, db mongo.Database, router *gin.Engine) {
	publicRoute := router.Group("/auth")
	initPublicUserRoutes(env, timeout, db, publicRoute)

	protectedRoute := router.Group("/", middlewares.JwtAuthMiddleWare(env.AccessTokenSecret))
	initProtectedCompanyRoute(env, timeout, db, protectedRoute.Group("companies"))
	initProtectedBankRoute(env, timeout, db, protectedRoute.Group("bank-services"))
	initProtectedTransactionRoute(env, timeout, db, protectedRoute.Group("transactions"))
	initProtectedCardRoute(env, timeout, db, protectedRoute.Group("cards"))
	initProtectedLoanRoute(env, timeout, db, protectedRoute.Group("active-loans"))
	initProtectedUserRoutes(env, timeout, db, protectedRoute.Group("user"))
}

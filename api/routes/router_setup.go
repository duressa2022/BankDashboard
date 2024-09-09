package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/config"
	"working.com/bank_dash/package/mongo"
)

func SetUp(env *config.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRoute := gin.Group("")
	SignUpRoute(env, timeout, db, publicRoute)

}

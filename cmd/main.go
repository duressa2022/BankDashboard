package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/api/routes"
	"working.com/bank_dash/config"
)

func main() {
	app, err := config.App()
	if err != nil {
		fmt.Println(err)
		return
	}

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()
	routes.SetUpRoute(env, timeout, db, router)
	router.Run(env.ServerAddress)
}

package main

import (
	"fmt"

	"working.com/bank_dash/config"
)


func main() {
	env := config.NewEnv();
	fmt.Println(env.AppEnv)
}
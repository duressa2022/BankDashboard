package config

import (
	"log"

	"github.com/spf13/viper"
)

// type for working all env varible
type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

// method for creating new Env
func NewEnv() *Env {
	env := Env{}
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("../")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cant find the docs to read", err)
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("error while loading the docs", err)
	}
	if env.AppEnv == "developemnt" {
		log.Println("app if running in dev env")
	}
	return &env

}

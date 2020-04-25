package config

import (
	"github.com/ditojanelidze/jwt-auth/config/initializers"
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	LoadEnvVariables()
	initializers.InitServer()
	initializers.DbInit()
	initializers.RedisInit()
}

func LoadEnvVariables() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalln("Unable to load environment varaibles\n", err)
	}
}
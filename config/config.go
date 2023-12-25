package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvVariables struct {
	HTTP_LISTEN_ADDRESS string `env:"HTTP_LISTEN_ADDRESS,required"`
	MONGO_DB_URL        string `env:"MONGO_DB_URL,required"`
	MONGO_DB_NAME       string `env:"MONGO_DB_NAME,required"`
	JWT_SECRET_KEY      string `env:"JWT_SECRET_KEY,required"`
}

var Env EnvVariables

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("unable to load .env file: %v", err)
	}
	if err := env.Parse(&Env); err != nil {
		log.Fatalf("unable to parse environment variables: %v", err)
	}
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	AppEnv    string
	DBUrl     string
	JWTSecret string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	cfg := &Config{
		Port:   getEnv("PORT", "8081"),
		AppEnv: getEnv("App_Env", "development"),
		DBUrl:  getEnv("DB_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}
	return val
}

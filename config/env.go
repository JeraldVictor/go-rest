package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()

	return Config{
		PublicHost: getEnv("SERVER_HOST", "http://localhost"),
		Port:       getEnv("SERVER_PORT", ":8080"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "root@123"),
		DBName:     getEnv("DB_NAME", "go_test"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

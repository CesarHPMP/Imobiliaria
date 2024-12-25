package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration values.
type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTsecret  string
}

// LoadConfig loads configuration from environment variables or .env file.
func LoadConfig() Config {
	// Load .env file if it exists
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return Config{
		Port:       GetEnv("PORT", "8080"),
		DBHost:     GetEnv("DB_HOST", "localhost"),
		DBPort:     GetEnv("DB_PORT", "5433"),
		DBUser:     GetEnv("DB_USER", "postgres"),
		DBPassword: GetEnv("DB_PASSWORD", "password"),
		DBName:     GetEnv("DB_NAME", "imobiliaria"),
		JWTsecret:  GetEnv("JWT_SECRET", "JWTsecret"),
	}
}

// GetEnv retrieves environment variables or returns a default value.
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

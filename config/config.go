package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		return nil, err
	}

	config := &Config{
		Port:       getEnv("PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "portDb"),
		DBUser:     getEnv("DB_USER", "youruser"),
		DBPassword: getEnv("DB_PASSWORD", "yourpassword"),
		DBName:     getEnv("DB_NAME", "yourdatabase"),
	}

	return config, nil
}

// Helper function to read an environment variable or return a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

package config

import (
	"os"
	"log"
)

type AppConfig struct {
	Port        string
	Token       string
}

var NewApp AppConfig

func SetConfig() {
	NewApp = AppConfig{
		Port:      getEnv("port", ":5000"),
		Token:getEnv("token", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else if defaultVal == "" {
		log.Fatalf("Environment variable %s cannot have a nil value", key)
	}
	return defaultVal
}
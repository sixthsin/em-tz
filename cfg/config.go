package cfg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db     DbConfig
	Server ServerConfig
}

type DbConfig struct {
	PostgresURL string
}

type ServerConfig struct {
	Port string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		Db: DbConfig{
			PostgresURL: os.Getenv("POSTGRES_URL"),
		},
		Server: ServerConfig{
			Port: os.Getenv("PORT"),
		},
	}
}

package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host           string
	Port           int
	User           string
	Password       string
	Name           string
	MigrationsPath string
}

type ServerConfig struct {
	Address string
}

func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	conf := &Config{
		Database: DatabaseConfig{
			Host:           os.Getenv("DB_HOST"),
			Port:           dbPort,
			User:           os.Getenv("DB_USER"),
			Password:       os.Getenv("DB_PASSWORD"),
			Name:           os.Getenv("DB_NAME"),
			MigrationsPath: os.Getenv("DB_MIGRATIONS_PATH"),
		},
		Server: ServerConfig{
			Address: os.Getenv("SERVER_ADDRESS"),
		},
	}

	return conf, nil
}

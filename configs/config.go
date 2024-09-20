package configs

import (
	"fmt"
	"os"
)

type Config struct {
	WebServerPort string
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	MigrationURL  string
	DBSource      string
}

func LoadConfig() (*Config, error) {
    cfg := &Config{
        WebServerPort: os.Getenv("WEB_SERVER_PORT"),
        DBDriver:      os.Getenv("DB_DRIVER"),
        DBHost:        os.Getenv("DB_HOST"),
        DBPort:        os.Getenv("DB_PORT"),
        DBUser:        os.Getenv("DB_USER"),
        DBPassword:    os.Getenv("DB_PASSWORD"),
        DBName:        os.Getenv("DB_NAME"),
        MigrationURL:  os.Getenv("MIGRATION_URL"),
        DBSource:      os.Getenv("DB_SOURCE"),
    }
    if cfg.WebServerPort == "" || cfg.DBDriver == "" || cfg.DBHost == "" ||
        cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" {
        return nil, fmt.Errorf("some envs are missing")
    }

    return cfg, nil
}
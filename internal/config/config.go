package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DBServer      string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	MaxUploadSize int64
	BatchSize     int
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	maxUploadSize, _ := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "33554432"), 10, 64)
	batchSize, _ := strconv.Atoi(getEnv("BATCH_SIZE", "1000"))

	return &Config{
		Port:          getEnv("PORT", "8080"),
		DBServer:      getEnv("DB_SERVER", "localhost"),
		DBPort:        getEnv("DB_PORT", "1433"),
		DBUser:        getEnv("DB_USER", "sa"),
		DBPassword:    getEnv("DB_PASSWORD", ""),
		DBName:        getEnv("DB_NAME", "CSVUploadDB"),
		MaxUploadSize: maxUploadSize,
		BatchSize:     batchSize,
	}
}

func InitDB(cfg *Config) (*sql.DB, error) {
	connString := fmt.Sprintf(
		"server=%s;port=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		cfg.DBServer,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("mssql", connString)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("Successfully connected to SQL Server")
	return db, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

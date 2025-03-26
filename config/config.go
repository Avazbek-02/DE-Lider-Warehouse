package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Postgres struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	JWT struct {
		Secret string
	}
	MinioUrl        string
	MinioUser       string
	MinIOSecredKey  string
	MinIOBucketName string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	// Postgres config
	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Postgres.DBName = os.Getenv("POSTGRES_DB")
	cfg.Postgres.SSLMode = os.Getenv("POSTGRES_SSLMODE")

	// JWT config
	cfg.JWT.Secret = os.Getenv("JWT_SECRET")

	// MinIO config
	cfg.MinioUrl = os.Getenv("MINIO_URL")
	cfg.MinioUser = os.Getenv("MINIO_USER")
	cfg.MinIOSecredKey = os.Getenv("MINIO_SECRET_KEY")
	cfg.MinIOBucketName = os.Getenv("MINIO_BUCKET_NAME")

	return cfg, nil
}

// Error codes
const (
	ErrorBadRequest     = "BAD_REQUEST"
	ErrorUnauthorized   = "UNAUTHORIZED"
	ErrorForbidden      = "FORBIDDEN"
	ErrorNotFound       = "NOT_FOUND"
	ErrorDuplicateKey   = "DUPLICATE_KEY"
	ErrorConflict       = "CONFLICT"
	ErrorInvalidRequest = "INVALID_REQUEST"
	ErrorInternalServer = "INTERNAL_SERVER_ERROR"
	ErrorInvalidPass    = "INVALID_PASSWORD"
)

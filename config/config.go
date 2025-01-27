package config

import (
	"api/internal/service/auth"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	ApiTimeout  time.Duration
	AppPort     int    `env:"APP_PORT"`
	DBUser      string `env:"POSTGRES_USER"`
	DBPassword  string `env:"POSTGRES_PASSWORD"`
	DBHost      string `env:"POSTGRES_HOST"`
	DBPort      string `env:"POSTGRES_PORT"`
	DBName      string `env:"POSTGRES_DB"`
	TokenConfig auth.TokenConfig
}

func MustLoadConfig() *Config {
	if err := godotenv.Load("/app/.env"); err != nil {
		panic("Warning: .env file not found, loading defaults or environment variables")
	}
	var cfg Config
	apiTimeout, err := time.ParseDuration(os.Getenv("API_TIMEOUT"))
	if err != nil {
		apiTimeout = 5 * time.Second
	}
	cfg.AppPort, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		cfg.AppPort = 8080
	}
	cfg.DBHost = os.Getenv("POSTGRES_HOST")
	cfg.DBPort = os.Getenv("POSTGRES_PORT")
	cfg.DBUser = os.Getenv("POSTGRES_USER")
	cfg.DBPassword = os.Getenv("POSTGRES_PASSWORD")
	cfg.DBName = os.Getenv("POSTGRES_DB")
	cfg.ApiTimeout = apiTimeout

	secret := []byte(os.Getenv("SECRET"))
	tokenTTL, err := time.ParseDuration(os.Getenv("TOKEN_TTL"))
	if err != nil {
		tokenTTL = 3 * 24 * time.Hour
	}
	refreshTokenTTL, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_TTL"))
	if err != nil {
		refreshTokenTTL = 20 * 24 * time.Hour
	}
	cfg.TokenConfig = auth.TokenConfig{
		Secret:          secret,
		TokenTTL:        tokenTTL,
		RefreshTokenTTL: refreshTokenTTL,
	}

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}

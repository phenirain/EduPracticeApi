package config

import (
	"api/internal/service/auth"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	ApiTimeout  time.Duration `env:"api_timeout"`
	AppPort     int           `env:"APP_PORT"`
	DBUser      string        `env:"POSTGRES_USER"`
	DBPassword  string        `env:"POSTGRES_PASSWORD"`
	DBHost      string        `env:"POSTGRES_HOST"`
	DBPort      string        `env:"POSTGRES_PORT"`
	DBName      string        `env:"POSTGRES_DB"`
	TokenConfig auth.TokenConfig
}

func MustLoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic("Warning: .env file not found, loading defaults or environment variables")
	}
	var cfg Config
	apiTimeout, err := time.ParseDuration(os.Getenv("API_TIMEOUT"))
	if err != nil {
		apiTimeout = 5 * time.Second
	}
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

package auth

import (
	"api/internal/domain/employees"
	"context"
	"time"
)

type AuthRepository interface {
	GetByLogin(ctx context.Context, login string) (*employees.Employee, error)
}

type UnifiedResponse struct {
	AuthResponse *AuthResponse `json:"auth_response,omitempty"`
	JwtResponse  *JwtResponse  `json:"jwt_response,omitempty"`
}

type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type JwtResponse struct {
	*AuthResponse
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role"`
}

type TokenConfig struct {
	Secret          []byte
	TokenTTL        time.Duration
	RefreshTokenTTL time.Duration
}

type AuthService struct {
	repo   AuthRepository
	config TokenConfig
}

func NewAuthService(repo AuthRepository, config TokenConfig) *AuthService {
	return &AuthService{repo, config}
}

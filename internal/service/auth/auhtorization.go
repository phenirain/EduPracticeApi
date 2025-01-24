package auth

import (
	"api/internal/lib/jwt"
	"context"
	"database/sql"
)

func (s *AuthService) TryLogin(ctx context.Context, request *AuthRequest) (*UnifiedResponse, error) {
	employee, err := s.repo.GetByLogin(ctx, request.Login)
	if err == sql.ErrNoRows {
		return &UnifiedResponse{
			AuthResponse: &AuthResponse{
				Status:  "Invalid credentials",
				Message: "No such login",
			},
		}, nil
	} else if err != nil {
		return nil, err
	}

	success := employee.CheckPassword(request.Password)
	if success {
		return &UnifiedResponse{
			JwtResponse: &JwtResponse{
				AuthResponse: &AuthResponse{
					Status:  "Success",
					Message: "Login successful",
				},
				Token:        jwt.MustGenerateToken(employee.Id, employee.Role.Name, s.config.Secret, s.config.tokenTTL),
				RefreshToken: jwt.MustGenerateToken(employee.Id, employee.Role.Name, s.config.Secret, s.config.refreshTokenTTL),
				Role:         employee.Role.Name,
			},
		}, nil
	} else {
		return &UnifiedResponse{
			AuthResponse: &AuthResponse{
				Status:  "Invalid credentials",
				Message: "Incorrect password",
			},
		}, nil
	}
}

package auth

import (
	"context"
	"database/sql"
)

func (s *AuthService) TryToLogin(ctx context.Context, request *AuthRequest) (*UnifiedResponse, error) {
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
				Token:        generateToken(employee.Id, employee.Role.Name),
				RefreshToken: generateRefreshToken(),
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

func VerifyToken(ctx context.Context, request *AuthRequest) (*UnifiedResponse, error) {
	// Verify JWT token here
	return nil, nil
}

func generateToken(userId int32, role string) string {
	// Generate JWT token here
	return "generated_jwt_token"
}

func generateRefreshToken() string {
	// Generate refresh token here
	return "generated_refresh_token"
}

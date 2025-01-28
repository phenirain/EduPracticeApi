package auth

import (
	"context"
	"database/sql"
)

func (s *AuthService) TryLogin(ctx context.Context, request *AuthRequest) (*AuthResponse, error) {
	employee, err := s.repo.GetByLogin(ctx, request.Login)
	if err == sql.ErrNoRows {
		return &AuthResponse{
			Status:  "Invalid credentials",
			Message: "No such login",
		}, nil
	} else if err != nil {
		return nil, err
	}
	
	success := employee.CheckPassword(request.Password)
	if success {
		return &AuthResponse{
			Status:  "Success",
			Message: "Login successful",
			Role:    employee.Role.Name,
		}, nil
	} else {
		return &AuthResponse{
			Status:  "Invalid credentials",
			Message: "Incorrect password",
		}, nil
	}
}

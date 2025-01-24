package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

func MustGenerateToken(userId int32, role string, secret []byte, duration time.Duration) string {
	claims := jwt.MapClaims{
		"id":   userId,
		"role": role,
		"exp":  time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(secret)
	return tokenString
}

func VerifyToken(tokenString string, secret []byte) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return false, fmt.Errorf("token parse error: %s", err.Error())
	}
	if !token.Valid {
		return false, ErrInvalidToken
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, errors.New("invalid token claims format")
	}
	_, ok = claims["id"]
	if !ok {
		return false, errors.New("can't get id from claims")
	}
	return true, nil
}

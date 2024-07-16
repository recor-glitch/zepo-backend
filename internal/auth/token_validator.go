package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

// ValidateToken validates a JWT token and returns the claims
func ValidateToken(tokenString string, secret []byte) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// ValidateAccessToken validates an access token and returns the claims
func ValidateAccessToken(tokenString string) (*Claims, error) {
	return ValidateToken(tokenString, AccessSecret)
}

// ValidateRefreshToken validates a refresh token and returns the claims
func ValidateRefreshToken(tokenString string) (*Claims, error) {
	return ValidateToken(tokenString, RefreshSecret)
}

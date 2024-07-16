package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateAccessToken generates a new JWT access token
func GenerateAccessToken(userID string) (string, error) {
	expirationTime := time.Now().Add(time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(AccessSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// GenerateRefreshToken generates a new JWT refresh token
func GenerateRefreshToken(userID string) (string, error) {
	// REFRESH TOKEN VALID FOR 180(6 MONTHS)/DAY
	expirationTime := time.Now().Add(180 * 24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(RefreshSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

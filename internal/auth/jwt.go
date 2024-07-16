package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var (
	AccessSecret  = []byte(os.Getenv("ACCESS_SECRET_KEY"))  // Replace with your access secret key
	RefreshSecret = []byte(os.Getenv("REFRESH_SECRET_KEY")) // Replace with your refresh secret key
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

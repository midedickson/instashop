package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/midedickson/instashop/config"
)

type JWTPayload map[string]any

type TokenGenOptions struct {
	Payload    JWTPayload
	ExpiryDate time.Time
}

type JWTClaim struct {
	Payload JWTPayload `json:"payload"`
	jwt.RegisteredClaims
}

func GenerateHS256Token(tokenGenOptions *TokenGenOptions) (string, error) {
	secret := config.GetJwtSecret()
	claims := &JWTClaim{
		Payload: tokenGenOptions.Payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: tokenGenOptions.ExpiryDate},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

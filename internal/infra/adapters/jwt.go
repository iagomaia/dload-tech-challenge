package adapters

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	jwtprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/jwt"
)

var (
	_ jwtprotocols.IJwt = (*JwtAdapter)(nil)
)

type JwtAdapter struct{}

func (a *JwtAdapter) Generate(userId string, claims map[string]any) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	jwtClaims := jwt.MapClaims{}
	for k, v := range claims {
		jwtClaims[k] = v
	}
	jwtClaims["exp"] = time.Now().Add(120 * time.Minute).Unix()
	jwtClaims["sub"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString([]byte(secret))
}

func (a *JwtAdapter) Verify(userToken string) (map[string]any, error) {
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := os.Getenv("JWT_SECRET")
		hmacSampleSecret := []byte(secret)
		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if mapClaims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if !mapClaims.VerifyExpiresAt(time.Now().Unix(), true) {
			return nil, errors.New("token expired")
		}
		claims := map[string]any{}
		for k, v := range mapClaims {
			claims[k] = v
		}
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

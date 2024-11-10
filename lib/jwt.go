package lib

import (
	"fmt"
	"time"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJWT(id uuid.UUID, name, email, role string) (string, error) {
	jwtToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id": id,
			"name": name,
			"email": email,
			"role": role,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	token, err := jwtToken.SignedString([]byte(config.Config.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyJWT(token string) (jwt.MapClaims, error) {
	tokenData, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWT_SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if !tokenData.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := tokenData.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
package lib

import (
	"fmt"
	"time"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJWT(id uuid.UUID) (string, error) {
	jwtToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id": id,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	token, err := jwtToken.SignedString([]byte(config.Config.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyJWT(token string) error {
	tokenStr, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWT_SECRET_KEY), nil
	})

	if err != nil {
		return err
	}

	if !tokenStr.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
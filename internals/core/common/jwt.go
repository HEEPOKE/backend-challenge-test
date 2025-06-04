package common

import (
	"time"

	"github.com/HEEPOKE/backend-challenge-test/pkg/configs"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(configs.Cfg.JWT_SECRET_KEY)
}

package utils

import (
	structures "client-service-go/config/configStruct"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	TokenTTL = 12 * time.Hour
)

type TokenClaims struct {
	jwt.StandardClaims
	ClientId int `json:"id"`
}

func ParseClientToken(accessToken string, appData *structures.AppData) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(appData.TokenConfig.ApiSecretClient), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenCliams")
	}

	return claims.ClientId, nil
}

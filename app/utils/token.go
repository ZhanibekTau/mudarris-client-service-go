package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	structures "user-service-go/config/configStruct"
)

const (
	TokenTTL = 12 * time.Hour
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"id"`
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

	return claims.UserId, nil
}

func ParseUstazToken(accessToken string, appData *structures.AppData) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(appData.TokenConfig.ApiSecretUstaz), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenCliams")
	}

	return claims.UserId, nil
}

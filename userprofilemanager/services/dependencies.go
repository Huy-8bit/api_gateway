package services

import (
	"time"
	"userprofilemanager/core"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateAuthToken(tokenString string) (string, time.Time, string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(core.GetEnvrionment("SECRET_KEY")), nil
		},
	)
	if err != nil {
		return "", time.Time{}, "", err
	}
	return claims["id"].(string), time.Unix(int64(claims["exp"].(float64)), 0), claims["typ"].(string), nil
}

func ValidateAccessToken(tokenString string) (string, time.Time, string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(core.GetEnvrionment("SECRET_KEY")), nil
		},
	)
	if err != nil {
		return "", time.Time{}, "", err
	}
	return claims["id"].(string), time.Unix(int64(claims["exp"].(float64)), 0), claims["typ"].(string), nil
}

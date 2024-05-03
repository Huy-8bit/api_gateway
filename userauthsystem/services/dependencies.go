package services

import (
	"time"
	"userauthsystem/core"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAcctoken(id string) string {
	option := jwt.SigningMethodHS256
	token := jwt.NewWithClaims(option, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"typ": "access",
	})
	var key []byte = []byte(core.GetEnvrionment("SECRET_KEY"))
	tokenString, _ := token.SignedString(key)
	return tokenString

}

func GetDataFromAccessToken(tokenString string) (interface{}, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(core.GetEnvrionment("SECRET_KEY")), nil
		},
	)
	if err != nil {
		return nil, err
	}

	return claims, nil

}

func ValidateAccessToken(tokenString string) (string, error) {
	var id string = core.RedisGet(tokenString)
	return id, nil
}

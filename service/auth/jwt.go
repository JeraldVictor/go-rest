package auth

import (
	"rest/config"
	"rest/types"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(user *types.User) (string, error) {
	secret := []byte(config.Envs.JWTSecret)
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSec)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(user.ID),
		"email":     user.Email,
		"user_name": user.Username,
		"expiresAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

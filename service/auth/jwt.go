package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"rest/config"
	"rest/types"
	"rest/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIdKey contextKey = "userId"

type JwtClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func CreateJWT(user *types.User) (string, error) {
	secret := []byte(config.Envs.JWTSecret)
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSec)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(user.ID),
		"email":     user.Email,
		"user_name": user.Username,
		"exp":       time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.Userstore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := getTokenFromHeader(r)
		if err != nil {
			utils.WriteError(w, http.StatusForbidden, err)
			return
		}

		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(*JwtClaims)
		userID, err := strconv.Atoi(claims.UserId)
		if err != nil {
			log.Printf("failed to convert userID to int: %v", err)
			permissionDenied(w)
			return
		}

		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		// Add the user to the context
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIdKey, u.ID)
		r = r.WithContext(ctx)

		// Call the function if the token is valid
		handlerFunc(w, r)
	}
}

func getTokenFromHeader(r *http.Request) (string, error) {
	authToken := r.Header.Get("Authorization")
	if authToken != "" {
		return authToken, nil
	}
	return "", fmt.Errorf("missing Authorization in header")
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserIdKey).(int)
	if !ok {
		return -1
	}

	return userID
}

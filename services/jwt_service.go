package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var signingKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID uint, expiresIn time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiresIn).Unix()
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": expirationTime,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func ParseToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		return 0, errors.New("token expired")
	}

	userID, ok := claims["sub"].(float64)
	if !ok || userID == 0 {
		return 0, errors.New("invalid subject")
	}

	return uint(userID), nil
}

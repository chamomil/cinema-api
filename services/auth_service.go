package services

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Signup(email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	return CreateUser(email, hash)
}

func Login(email, password string) (tokenString string, err error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	return GenerateToken(user.ID, time.Hour*24*3)
}

package services

import (
	"cinema/db"
	"cinema/models"
)

func GetUserByID(userID uint) (*models.User, error) {
	var user *models.User
	err := db.DB.First(&user, userID).Error
	return user, err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func CreateUser(email string, hash []byte) error {
	user := models.User{Email: email, Password: string(hash)}
	return db.DB.Create(&user).Error
}

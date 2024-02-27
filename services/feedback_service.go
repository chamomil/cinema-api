package services

import (
	"cinema/db"
	"cinema/models"
)

func CreateFeedback(userID uint, message string) error {
	feedback := models.Feedback{
		UserID:  userID,
		Message: message,
	}
	return db.DB.Create(&feedback).Error
}

func GetFeedbacksByUserID(userID uint) ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	if err := db.DB.Where("user_id = ?", userID).Find(&feedbacks).Error; err != nil {
		return nil, err
	}
	return feedbacks, nil
}

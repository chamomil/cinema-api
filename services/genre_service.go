package services

import (
	"cinema/db"
	"cinema/models"
)

func FindGenres() []models.Genre {
	var genres []models.Genre
	db.DB.Find(&genres)
	return genres
}

func FindGenreByID(id uint) (models.Genre, error) {
	var genre models.Genre
	err := db.DB.First(&genre, id).Error
	return genre, err
}

func CreateGenre(input models.CreateGenreInput) (models.Genre, error) {
	genre := models.Genre{Name: input.Name}
	err := db.DB.Create(&genre).Error
	return genre, err
}

func UpdateGenre(id uint, input models.UpdateGenreInput) (models.Genre, error) {
	var genre models.Genre
	if err := db.DB.First(&genre, id).Error; err != nil {
		return genre, err
	}

	err := db.DB.Model(&genre).Updates(input).Error
	return genre, err
}

func DeleteGenre(id uint) error {
	var genre models.Genre
	if err := db.DB.First(&genre, id).Error; err != nil {
		return err
	}
	if err := db.DB.Delete(&genre).Error; err != nil {
		return err
	}
	return nil
}

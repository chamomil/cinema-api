package services

import (
	"cinema/db"
	"cinema/models"
)

func GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie
	if err := db.DB.Preload("Genres").Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func GetMovieByID(id uint) (models.Movie, error) {
	var movie models.Movie
	if err := db.DB.Preload("Genres").First(&movie, id).Error; err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func CreateMovie(input models.CreateMovieInput) (models.Movie, error) {
	movie := models.Movie{
		Title:       input.Title,
		Description: input.Description,
		Duration:    input.Duration,
		Image:       input.Image,
	}
	if err := db.DB.Create(&movie).Error; err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func UpdateMovie(id uint, input models.UpdateMovieInput) (models.Movie, error) {
	var movie models.Movie
	if err := db.DB.First(&movie, id).Error; err != nil {
		return models.Movie{}, err
	}
	if err := db.DB.Model(&movie).Updates(input).Error; err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func DeleteMovie(id uint) error {
	var movie models.Movie
	if err := db.DB.First(&movie, id).Error; err != nil {
		return err
	}
	if err := db.DB.Delete(&movie).Error; err != nil {
		return err
	}
	return nil
}

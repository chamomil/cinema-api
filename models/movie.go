package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string
	Description string
	Duration    int
	Genres      []Genre `gorm:"many2many:movie_genres"`
	Image       string
}

type CreateMovieInput struct {
	Title       string
	Description string
	Duration    int
	Genres      []int
	Image       string
}

type UpdateMovieInput struct {
	Title       string
	Description string
	Duration    int
	Genres      []int
	Image       string
}

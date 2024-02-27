package models

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name string `gorm:"type:text"`
}

type CreateGenreInput struct {
	Name string
}

type UpdateGenreInput struct {
	Name string
}

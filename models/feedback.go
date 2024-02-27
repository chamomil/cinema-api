package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	UserID  uint
	Message string
}

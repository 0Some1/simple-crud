package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	name string
	text string
}

package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Name string `json:"name"`
	Text string `json:"text"`
}

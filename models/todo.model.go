package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title string `json:"title" validate:"required,min=3,max=30"`
	Body  string `json:"body" validate:"required"`
	Done  bool   `json:"done"`
}

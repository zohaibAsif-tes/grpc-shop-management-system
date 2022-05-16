package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	Name  string
	Total float32
}

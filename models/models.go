package models

import "gorm.io/gorm"

type Author struct {
	*gorm.Model
	Name string
}

type Book struct {
	*gorm.Model
	Name   string
	Author string
}

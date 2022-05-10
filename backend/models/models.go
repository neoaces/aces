package models

// Here, all the typing interfaces for the sqlite database are defined.

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Name   string
	Answer string
}

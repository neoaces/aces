package models

// Here, all the typing interfaces for the sqlite database are defined.

import "gorm.io/gorm"

type User struct {
	gorm.Model
	username string
	cardsets []CardSet
}

type CardSet struct {
	gorm.Model
	name  string
	class string
	cards []Card `gorm:"foreignKey:Cardset"`
}

type Card struct {
	gorm.Model
	Name   string
	Answer string
}

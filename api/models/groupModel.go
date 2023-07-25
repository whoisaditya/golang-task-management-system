package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Users []User `gorm:"many2many:user_groups;"`
}

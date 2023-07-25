package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`

	Password string    `gorm:"not null"`
	Phone    string    
	DOB      time.Time 

	// UserRole uint  `gorm:"default:0"` // 0: user, 1: admin, 2:owner
	// Role     Role `gorm:"foreignKey:UserRole;references:id"`
}

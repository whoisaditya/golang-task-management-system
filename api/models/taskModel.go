package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`

	CreatedBy uint  `gorm:"not null;default:0"`
	User      User `gorm:"foreignKey:CreatedBy;references:id"`

	PlannedStartTime time.Time
	PlannedEndTime   time.Time 
	ActualStartTime  time.Time 
	ActualEndTime    time.Time 
	Seconds            int64

	Status           int     `gorm:"default:0"` // 0 = new, 1 = ongoing, 2 = completed
	AssignedToGroups []Group `gorm:"many2many:task_assigned_to_group"`
}

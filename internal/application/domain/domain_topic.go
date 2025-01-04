package domain

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	ID      int    `gorm:"primarykey"`
	Title   string `gorm:"type:varchar(255); not null"`
	Content string `gorm:"type:varchar(255); not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// relation
	UserID int
	GoalID int
}

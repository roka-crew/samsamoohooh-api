package domain

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	ID        int       `gorm:"primarykey"`
	Deadline  time.Time `gorm:"type:date; not null"`
	PageRange int       `gorm:"type:int; default:0; not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// relation
	GroupID int
	Topics  []Topic
}

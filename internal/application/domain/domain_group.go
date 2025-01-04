package domain

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	ID               int     `gorm:"primarykey"`
	BookTitle        string  `gorm:"type:varchar(255); not null"`
	BookAuthor       string  `gorm:"type:varchar(255); not null"`
	BookPageMax      int     `gorm:"type:int; default:0; not null"`
	BookPageCount    int     `gorm:"type:int; default:0; not null"`
	BookPublisher    *string `gorm:"type:varchar(255)"`
	BookIntroduction *string `gorm:"type:varchar(255)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// relation
	Users []User `gorm:"many2many:user_group;"`
	Goals []Goal
}

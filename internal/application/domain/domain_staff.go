package domain

import (
	"time"

	"gorm.io/gorm"
)

type Staff struct {
	ID        int    `gorm:"primarykey"`
	AccountID string `gorm:"varchar(255); not null"`
	Password  string `gorm:"varchar(255); not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

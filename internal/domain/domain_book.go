package domain

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement"`
	Title     string         `gorm:"column:title;type:varchar(256);not null;index"`
	PageCount int            `gorm:"column:page_count;type:integer;not null"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`

	UserBooks  []UserBook  `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Milestones []Milestone `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

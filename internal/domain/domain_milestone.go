package domain

import (
	"time"

	"gorm.io/gorm"
)

type Milestone struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement"`
	BookID    uint           `gorm:"column:book_id;not null;index"` 
	Content   string         `gorm:"column:content;type:text;not null"`
	DueDate   time.Time      `gorm:"column:due_date"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`

	Book     *Book     `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Thoughts []Thought `gorm:"foreignKey:MilestoneID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

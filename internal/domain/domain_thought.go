package domain

import (
	"time"

	"gorm.io/gorm"
)

type Thought struct {
	ID          uint           `gorm:"column:id;primaryKey;autoIncrement"`
	MilestoneID uint           `gorm:"column:milestone_id;not null;index"`
	UserID      uint           `gorm:"column:user_id;not null;index"`
	Content     string         `gorm:"column:content;type:text;not null"`
	IsAnonymous bool           `gorm:"column:is_anonymous;type:boolean;not null;default:true"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`

	Milestone *Milestone `gorm:"foreignKey:MilestoneID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      *User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ThoughtRepository interface {
}

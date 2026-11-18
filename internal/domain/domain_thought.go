package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Thoughts []Thought

type Thought struct {
	ID          uint           `gorm:"column:id;primaryKey;autoIncrement"`
	UserID      uint           `gorm:"column:user_id;not null;index"`
	MilestoneID uint           `gorm:"column:milestone_id;not null;index"`
	Content     string         `gorm:"column:content;type:text;not null"`
	IsAnonymous bool           `gorm:"column:is_anonymous;type:boolean;not null;default:true"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`

	Milestone *Milestone `gorm:"foreignKey:MilestoneID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      *User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ThoughtRepository interface {
	CreateThought(context.Context, CreateThoughtParam) (*Thought, error)
	ListThought(context.Context, ListThoughtParam) (*Thought, error)
	ListThoughts(context.Context, ListThoughtsParam) (Thoughts, error)
	PatchThought(context.Context, PatchThoughtParam) error
	PatchThoughts(context.Context, PatchThoughtsParam) error
	DeleteThought(context.Context, DeleteThoughtParam) error
	DeleteThoughts(context.Context, DeleteThoughtsParam) error
}

type CreateThoughtParam struct {
	UserID      uint
	MilestoneID uint
	Content     string
	IsAnonymous bool
}

type ListThoughtParam struct {
	ID          uint
	UserID      uint
	MilestoneID uint
	Content     *string
	IsAnonymous *bool

	WithUser     bool
	WithMileston bool
}

type ListThoughtsParam struct {
	// Filters
	IDs           []uint
	UserIDs       []uint
	MilestoneIDs  []uint
	Contents      []string
	IsAnonymouses []bool

	// Pagination & Sorting
	Limit   int
	Offset  int
	OrderBy []string // []string{"created_at DESC"}
}

type PatchThoughtParam struct {
	ID          uint
	UserID      *uint
	MilestoneID *uint
	Content     *string
	IsAnonymous *bool
}

type PatchThoughtsParam []PatchThoughtParam

type DeleteThoughtParam struct {
	ID           uint
	UserID       *uint
	MilestoneID  *uint
	IsHardDelete bool
}

type DeleteThoughtsParam []DeleteThoughtParam

package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Milestones []Milestone

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

type MilestoneRepository interface {
	CreateMilestone(context.Context, CreateMilestoneParam) (Milestone, error)
	ListMilestone(context.Context, ListMilestoneParam) (Milestone, error)
	ListMilestones(context.Context, ListMilestonesParam) (Milestones, error)
	PatchMilestone(context.Context, PatchMilestoneParam) error
	PatchMilestones(context.Context, PatchMilestonesParam) error
	DeleteMilestone(context.Context, DeleteMilestoneParam) error
	DeleteMilestones(context.Context, DeleteMilestonesParam) error
}

type CreateMilestoneParam struct {
	BookID  uint
	Content string
	DueDate time.Time
}

type ListMilestoneParam struct {
	ID     uint
	BookID *uint

	// Relations
	WithBook     bool
	WithThoughts bool
}

type ListMilestonesParam struct {
	// Filters
	IDs         []uint
	BookIDs     []uint
	Contents    []string
	DueDateFrom *time.Time
	DueDateTo   *time.Time

	// Relations
	WithBook     bool
	WithThoughts bool

	// Pagination & Sorting
	Limit   int
	Offset  int
	OrderBy []string // e.g., []string{"due_date ASC"}
}

type PatchMilestoneParam struct {
	ID      uint
	BookID  *uint
	Content *string
	DueDate *time.Time
}

type PatchMilestonesParam []PatchMilestoneParam

type DeleteMilestoneParam struct {
	ID           uint
	BookID       *uint
	IsHardDelete bool
}

type DeleteMilestonesParam []DeleteMilestoneParam

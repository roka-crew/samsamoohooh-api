package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Books []Book

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

type BookRepository interface {
	CreateBook(context.Context, CreateBookParam) (Book, error)
	ListBook(context.Context, ListBookParam) (Book, error)
	ListBooks(context.Context, ListBooksParam) (Books, error)
	PatchBook(context.Context, PatchBookParam) error
	PatchBooks(context.Context, PatchBooksParam) error
	DeleteBook(context.Context, DeleteBookParam) error
	DeleteBooks(context.Context, DeleteBooksParam) error
}

type CreateBookParam struct {
	Title     string
	PageCount int
}

type ListBookParam struct {
	ID        uint
	Title     *string
	PageCount uint

	// Relations
	WithUserBooks  bool
	WithMilestones bool
}

type ListBooksParam struct {
	// Filters
	IDs        []uint
	Titles     []string
	PageCounts []int

	// Relations
	WithUserBooks  bool
	WithMilestones bool

	// Pagination & Sorting
	Limit   int
	Offset  int
	OrderBy []string // e.g., []string{"created_at DESC"}
}

type PatchBookParam struct {
	ID        uint
	Title     *string
	PageCount *int
}

type PatchBooksParam []PatchBookParam

type DeleteBookParam struct {
	ID           uint
	Title        *string
	IsHardDelete bool
}

type DeleteBooksParam []DeleteBookParam

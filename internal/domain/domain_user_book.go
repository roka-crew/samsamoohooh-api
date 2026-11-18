package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserBooks []UserBook

type UserBook struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement"`
	UserID    uint           `gorm:"column:user_id;not null;index"`
	BookID    uint           `gorm:"column:book_id;not null;index"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`

	User *User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Book *Book `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserBookRepository interface {
	CreateUserBook(context.Context, CreateUserBookParam) (UserBook, error)
	CreateUserBooks(context.Context, CreateUserBooksParam) (UserBooks, error)
	ListUserBook(context.Context, ListUserBookParam) (UserBook, error)
	ListUserBooks(context.Context, ListUserBooksParam) (UserBooks, error)
	PatchUserBook(context.Context, PatchUserBookParam) error
	PatchUserBooks(context.Context, PatchUserBooksParam) error
	DeleteUserBook(context.Context, DeleteUserBookParam) error
	DeleteUserBooks(context.Context, DeleteUserBooksParam) error
}

type CreateUserBookParam struct {
	UserID uint
	BookID uint
}

type CreateUserBooksParam []CreateUserBookParam

type ListUserBookParam struct {
	ID     uint
	UserID *uint
	BookID *uint

	// Relations
	WithUser bool
	WithBook bool
}

type ListUserBooksParam struct {
	// Filters
	IDs     []uint
	UserIDs []uint
	BookIDs []uint

	// Relations
	WithUser bool
	WithBook bool

	// Pagination & Sorting
	Limit   int
	Offset  int
	OrderBy []string // e.g., []string{"created_at DESC"}
}

type PatchUserBookParam struct {
	ID     uint
	UserID *uint
	BookID *uint
}

type PatchUserBooksParam []PatchUserBookParam

type DeleteUserBookParam struct {
	ID           uint
	UserID       *uint
	BookID       *uint
	IsHardDelete bool
}

type DeleteUserBooksParam []DeleteUserBookParam

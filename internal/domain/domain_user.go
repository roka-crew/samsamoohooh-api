package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Users []User

type User struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string         `gorm:"column:name;type:varchar(100);not null"`
	DeviceID  string         `gorm:"column:device_id;type:varchar(255);not null;unique"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`

	UserBooks []UserBook `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Thoughts  []Thought  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserRepository interface {
	CreateUser(context.Context, CreateUserParam) (*User, error)
	ListUser(context.Context, ListUserParam) (*User, error)
	ListUsers(context.Context, ListUsersParam) (Users, error)
	PatchUser(context.Context, PatchUserParam) error
	PatchUsers(context.Context, PatchUsersParam) error
	DeleteUser(context.Context, DeleteUserParam) error
	DeleteUsers(context.Context, DeleteUsersParam) error
}

type CreateUserParam struct {
	Name     string
	DeviceID string
}

type ListUserParam struct {
	ID           uint
	WithBooks    bool
	WithThoughts bool
}

type ListUsersParam struct {
	// Filters
	IDs       []uint
	Name      []string
	DeviceIDs []string

	// Pagination & Sorting
	Limit   int
	Offset  int
	OrderBy []string // []string{"created_at DESC"}

	// Preloads
	WithUserBooks bool
	WithThoughts  bool
}

type PatchUserParam struct {
	ID       uint `validate:"required"`
	Name     *string
	DeviceID *string
}

type PatchUsersParam []PatchUserParam

type DeleteUserParam struct {
	ID           uint
	DeviceID     *string
	IsHardDelete bool
}

type DeleteUsersParam []DeleteUserParam

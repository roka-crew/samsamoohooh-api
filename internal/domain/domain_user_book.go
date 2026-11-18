package domain

import (
	"time"

	"gorm.io/gorm"
)

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

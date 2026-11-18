package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement"`
	DeviceID  string         `gorm:"column:device_id;type:varchar(255);not null;unique"`
	Name      string         `gorm:"column:name;type:varchar(100);not null"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`

	UserBooks []UserBook `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Thoughts  []Thought  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
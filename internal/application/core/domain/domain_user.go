package domain

import (
	"time"

	"gorm.io/gorm"
)

type Provider string

const (
	ProviderGoogle Provider = "GOOGLE"
	ProviderApple  Provider = "APPLE"
	ProviderKakao  Provider = "KAKAKO"
)

type User struct {
	ID         int      `gorm:"primary_key"`
	Nickname   string   `gorm:"type:varchar(255); not null"`
	Resolution string   `gorm:"type:varchar(255); null"`
	Provider   Provider `gorm:"type:enum('GOOGLE', 'APPLE', 'KAKAO'); not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// relation
	Groups []*Group
	Topics []Topic
}

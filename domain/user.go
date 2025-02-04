package domain

import (
	"gorm.io/gorm"
)

type Provider string

const (
	ProviderGoogle Provider = "GOOGLE"
	ProviderApple  Provider = "APPLE"
	ProviderKakao  Provider = "KAKAO"
)

type User struct {
	gorm.Model
	Nickname   string   `gorm:"type:varchar(255); not null"`
	Resolution *string  `gorm:"type:varchar(255); null"`
	Provider   Provider `gorm:"type:enum('GOOGLE', 'APPLE', 'KAKAO'); not null"`

	// relation
	Groups []Group `gorm:"many2many:user_group;"`
	Topics []Topic
}

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Empty() bool {
	return u.Len() == 0
}

func (u Users) First() User {
	if u.Empty() {
		return User{}
	}
	return u[0]
}

func (u Users) Last() User {
	if u.Empty() {
		return User{}
	}
	return u[u.Len()-1]
}

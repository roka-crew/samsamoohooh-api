package domain

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	BookTitle        string  `gorm:"type:varchar(255); not null"`
	BookAuthor       string  `gorm:"type:varchar(255); not null"`
	BookPageMax      int     `gorm:"type:int; default:0; not null"`
	BookPageCount    int     `gorm:"type:int; default:0; not null"`
	BookPublisher    *string `gorm:"type:varchar(255)"`
	BookIntroduction *string `gorm:"type:varchar(255)"`

	// relation
	Users []User `gorm:"many2many:user_group;"`
	Goals []Goal
}

type Groups []Group

func (g Groups) Len() int {
	return len(g)
}

func (g Groups) Empty() bool {
	return g.Len() == 0
}

func (g Groups) First() Group {
	if g.Empty() {
		return Group{}
	}
	return g[0]
}

func (g Groups) Last() Group {
	if g.Empty() {
		return Group{}
	}
	return g[g.Len()-1]
}

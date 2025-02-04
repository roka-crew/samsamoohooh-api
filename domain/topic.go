package domain

import (
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255); not null"`
	Content string `gorm:"type:varchar(255); not null"`

	// relation
	UserID int
	GoalID int
}

type Topics []Topic

func (t Topics) Len() int {
	return len(t)
}

func (t Topics) Empty() bool {
	return t.Len() == 0
}

func (t Topics) First() Topic {
	if t.Empty() {
		return Topic{}
	}
	return t[0]
}

func (t Topics) Last() Topic {
	if t.Empty() {
		return Topic{}
	}
	return t[t.Len()-1]
}

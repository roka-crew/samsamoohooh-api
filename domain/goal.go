package domain

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	Deadline  time.Time `gorm:"type:date; not null"`
	PageRange int       `gorm:"type:int; default:0; not null"`

	// relation
	GroupID int
	Topics  []Topic
}

type Goals []Goal

func (g Goals) Len() int {
	return len(g)
}

func (g Goals) Empty() bool {
	return g.Len() == 0
}

func (g Goals) First() Goal {
	if g.Empty() {
		return Goal{}
	}
	return g[0]
}

func (g Goals) Last() Goal {
	if g.Empty() {
		return Goal{}
	}
	return g[g.Len()-1]
}

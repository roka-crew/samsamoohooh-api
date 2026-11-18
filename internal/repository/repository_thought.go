package repository

import (
	"github.com/roka-crew/samsamoohooh-api/internal/database/sqlite"
	"github.com/roka-crew/samsamoohooh-api/internal/domain"
)

type ThoughtRepository struct {
	db *sqlite.Sqlite
}

func NewThoughtRepository(
	db *sqlite.Sqlite,
) domain.ThoughtRepository {
	return &ThoughtRepository{
		db: db,
	}
}


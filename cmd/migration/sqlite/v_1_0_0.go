package main

import (
	"log"

	"github.com/roka-crew/samsamoohooh-api/internal/database/sqlite"
	"github.com/roka-crew/samsamoohooh-api/internal/domain"
	"github.com/roka-crew/samsamoohooh-api/pkg/config"
)

func main() {
	cfg, err := config.NewConfig("./configs/samsamoohooh.yaml")
	if err != nil {
		log.Panicf("failed to new config: %v", err)
	}

	sqliteDB, err := sqlite.NewSqlite(cfg)
	if err != nil {
		log.Panicf("failed to new sqlite: %v", err)
	}

	if err := sqliteDB.AutoMigrate(
		&domain.User{},
		&domain.Book{},
		&domain.Thought{},
		&domain.Milestone{},
		&domain.UserBook{},
	); err != nil {
		log.Panicf("failed to auto migrate: %v\n", err)
	}
}

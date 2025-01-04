package main

import (
	"log"
	"samsamoohooh-api/internal/infra/config"
	"samsamoohooh-api/internal/infra/persistence/mysql"
)

func main() {
	cfg, err := config.NewConfig("./configs/env.yaml")
	if err != nil {
		log.Panicf("failed new config, inspect: %v\n", err)
	}

	mysql, err := mysql.NewMysql(cfg)
	if err != nil {
		log.Panicf("failed new config, inspect: %v\n", err)
	}

	if err := mysql.Migrate(); err != nil {
		log.Panicf("failed to migrate, insepct: %v\n", err)
	}
}

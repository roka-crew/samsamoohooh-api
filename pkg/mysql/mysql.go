package mysql

import (
	"fmt"
	"samsamoohooh-api/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	format = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

type Mysql struct {
	*gorm.DB
}

func New(cfg config.Config) (*Mysql, error) {
	dsn := fmt.Sprintf(format,
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.Database,
	)

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			TranslateError: true,
		},
	)
	if err != nil {
		return nil, err
	}

	return &Mysql{DB: db}, nil
}

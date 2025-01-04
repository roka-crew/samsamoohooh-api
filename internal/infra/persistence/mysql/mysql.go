package mysql

import (
	"fmt"
	"samsamoohooh-api/internal/application/domain"
	"samsamoohooh-api/internal/infra/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	format = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

type Mysql struct {
	*gorm.DB
}

func NewMysql(
	config *config.Config,
) (*Mysql, error) {
	dsn := fmt.Sprintf(format,
		config.Mysql.User,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database,
	)

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			TranslateError: true,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open gorm: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to unwarp sql.DB: %w", err)
	}
	sqlDB.SetMaxIdleConns(config.MysqlConn.SetMaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MysqlConn.SetMaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.MysqlConn.SetConnMaxLifeTime)

	return &Mysql{DB: db}, nil
}

func (m *Mysql) Migrate() error {
	return m.DB.AutoMigrate(
		&domain.User{},
		&domain.Group{},
		&domain.Topic{},
		&domain.Goal{},
		&domain.Staff{},
	)
}

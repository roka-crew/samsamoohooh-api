package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Listen string `yaml:"listen"`
	} `yaml:"server"`

	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`

	MysqlConn struct {
		SetMaxIdleConns    int      `yaml:"setMaxIdleConns"`
		SetMaxOpenConns    int      `yaml:"setMaxOpenConns"`
		SetConnMaxLifeTime duration `yaml:"setConnMaxLifeTime"`
	} `yaml:"mysqlConn"`

	Token struct {
		AccessExp  duration `yaml:"accessExp"`
		RefreshExp duration `yaml:"refreshExp"`
		NotBefore  duration `yaml:"notBefore"`
		SecretKey  string   `yaml:"secretKey"`
	} `yaml:"token"`
}

func NewConfig(path string) (*Config, error) {
	config := Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode YAML file: %w", err)
	}

	return &config, nil
}

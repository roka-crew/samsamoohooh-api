package config

import (
	"fmt"
	"os"
	"time"

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

	MysqlConn MysqlConn `yaml:"mysqlConn"`

	Token Token `yaml:"token"`
}

type MysqlConn struct {
	SetMaxIdleConns    int           `yaml:"setMaxIdleConns"`
	SetMaxOpenConns    int           `yaml:"setMaxOpenConns"`
	SetConnMaxLifeTime time.Duration `yaml:"setConnMaxLifeTime"`
}

func (m *MysqlConn) UnmarshalYAML(value *yaml.Node) error {
	// 중간 구조체를 사용하여 YAML을 언마샬링
	type Alias struct {
		SetMaxIdleConns    int    `yaml:"setMaxIdleConns"`
		SetMaxOpenConns    int    `yaml:"setMaxOpenConns"`
		SetConnMaxLifeTime string `yaml:"setConnMaxLifeTime"`
	}

	aux := &Alias{}
	if err := value.Decode(aux); err != nil {
		return fmt.Errorf("failed to decode YAML node to ailas: %w", err)
	}

	duration, err := time.ParseDuration(aux.SetConnMaxLifeTime)
	if err != nil {
		return fmt.Errorf("invalid duration format for SetConnMaxLifeTime: %w", err)
	}

	// 값을 원본 구조체로 매핑
	m.SetMaxIdleConns = aux.SetMaxIdleConns
	m.SetMaxOpenConns = aux.SetMaxOpenConns
	m.SetConnMaxLifeTime = duration

	return nil
}

type Token struct {
	AccessExp  time.Duration `yaml:"accessExp"`
	RefreshExp time.Duration `yaml:"refreshExp"`
	NotBefore  time.Duration `yaml:"notBefore"`
	SecretKey  []byte        `yaml:"secretKey"`
}

func (t *Token) UnmarshalYAML(value *yaml.Node) error {
	// 중간 구조체를 사용하여 YAML을 언마샬링
	type Alias struct {
		AccessExp  string `yaml:"accessExp"`
		RefreshExp string `yaml:"refreshExp"`
		NotBefore  string `yaml:"notBefore"`
		SecretKey  string `yaml:"secretKey"`
	}

	aux := &Alias{}
	if err := value.Decode(aux); err != nil {
		return fmt.Errorf("failed to decode YAML node to alias: %w", err)
	}

	accessExp, err := time.ParseDuration(aux.AccessExp)
	if err != nil {
		return fmt.Errorf("invalid duration format for AccessExp: %w", err)
	}

	refreshExp, err := time.ParseDuration(aux.RefreshExp)
	if err != nil {
		return fmt.Errorf("invalid duration format for RefreshExp: %w", err)
	}

	notBefore, err := time.ParseDuration(aux.NotBefore)
	if err != nil {
		return fmt.Errorf("invalid duration format for NotBefore: %w", err)
	}

	// SecretKey는 문자열을 바이트 배열로 변환
	t.SecretKey = []byte(aux.SecretKey)
	t.AccessExp = accessExp
	t.RefreshExp = refreshExp
	t.NotBefore = notBefore

	return nil
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

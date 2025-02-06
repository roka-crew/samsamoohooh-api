package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Listen string `yaml:"listen"`

	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBname   string `yaml:"dbName"`
	} `yaml:"mysql"`

	JWT struct {
		SecretKey string `yaml:"secret"`
	} `yaml:"jwt"`
}

func New(path string) (*Config, error) {
	cfg := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

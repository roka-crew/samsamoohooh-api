package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {

	config, err := NewConfig("../../../configs/env.yaml")
	if err != nil {
		t.Fatalf("failed new config, inspect: %v\n", err)
	}

	// Server 설정 체크
	if config.Server.Addr == "" {
		t.Error("Server.Addr should not be empty")
	}

	// MySQL 설정 체크
	if config.Mysql.User == "" {
		t.Error("Mysql.User should not be empty")
	}
	if config.Mysql.Password == "" {
		t.Error("Mysql.Password should not be empty")
	}
	if config.Mysql.Host == "" {
		t.Error("Mysql.Host should not be empty")
	}
	if config.Mysql.Port == 0 {
		t.Error("Mysql.Port should not be zero")
	}
	if config.Mysql.Database == "" {
		t.Error("Mysql.Database should not be empty")
	}

	// MySQL Pool 설정 체크
	if config.MysqlConn.SetMaxIdleConns == 0 {
		t.Error("MysqlPool.SetMaxIdleConns should not be zero")
	}
	if config.MysqlConn.SetMaxOpenConns == 0 {
		t.Error("MysqlPool.SetMaxOpenConns should not be zero")
	}
	if config.MysqlConn.SetConnMaxLifeTime == 0 {
		t.Error("MysqlPool.SetConnMaxLifeTime should not be zero")
	}

	// Token 설정 체크
	if config.Token.AccessExp == 0 {
		t.Error("Token.AccessExp should not be zero")
	}
	if config.Token.RefreshExp == 0 {
		t.Error("Token.RefreshExp should not be zero")
	}
	if len(config.Token.SecretKey) == 0 {
		t.Error("Token.SecretKey should not be empty")
	}
}

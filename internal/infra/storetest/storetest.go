package storetest

import (
	"context"
	"fmt"
	"samsamoohooh-api/internal/infra/persistence/mysql"
	"samsamoohooh-api/internal/infra/validator"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DetafultCtx = context.Background()
var testcontainer testcontainers.Container

func GetMysql(t *testing.T) *mysql.Mysql {
	ctx := context.Background()
	host, err := testcontainer.Host(ctx)
	if err != nil {
		t.Errorf("컨테이너 호스트 정보 조회 실패: %v", err)
	}

	port, err := testcontainer.MappedPort(ctx, "3306")
	if err != nil {
		t.Errorf("컨테이너 포트 정보 조회 실패: %v", err)
	}

	dsn := fmt.Sprintf(
		"testuser:testpass@tcp(%s:%s)/testdb?charset=utf8mb4&parseTime=True&loc=Local",
		host,
		port.Port(),
	)

	db, err := gorm.Open(mysqldriver.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Errorf("데이터베이스 연결 실패: %v", err)
	}

	mysql := &mysql.Mysql{DB: db}
	if err := mysql.Migrate(); err != nil {
		t.Errorf("데이터베이스 Migrate 실패: %v", err)
	}

	return mysql
}

func GetValidator() *validator.Validator {
	return validator.NewValidator()
}

func SetUp() error {
	containerReq := testcontainers.ContainerRequest{
		Image:        "mysql:8.0",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "testpass",
			"MYSQL_DATABASE":      "testdb",
			"MYSQL_USER":          "testuser",
			"MYSQL_PASSWORD":      "testpass",
		},
		WaitingFor: wait.ForAll(
			wait.ForLog("port: 3306  MySQL Community Server - GPL"),
			wait.ForListeningPort("3306/tcp"),
		),
	}

	container, err := testcontainers.GenericContainer(DetafultCtx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerReq,
		Started:          true,
	})
	if err != nil {
		return fmt.Errorf("컨테이너 시작 실패: %v", err)
	}
	testcontainer = container

	return nil
}

func Shutdwon() error {
	return testcontainer.Terminate(DetafultCtx)
}

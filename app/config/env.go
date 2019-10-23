package config

import (
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"

	"os"
)

var Set = wire.NewSet(
	MysqlConfig,
)

func MysqlConfig() *mysql.Config {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASS")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DB_ADDR")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.ParseTime = true
	return cfg
}

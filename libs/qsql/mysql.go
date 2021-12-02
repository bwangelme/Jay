package qsql

import (
	"database/sql"
	"qae/helper"
	"qae/libs/conf"
	"qae/logger"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MySQLdb *sql.DB
)

func InitMySQL() {
	var err error

	cfg := conf.GetMySQLConfig()
	MySQLdb, err = sql.Open("mysql", cfg.DSN())
	helper.EnsureOK(err, map[string]interface{}{
		"msg": "init mysql failed",
		"dsn": cfg.DSN(),
	})
	logger.Infof("Init MySQL Success")
}

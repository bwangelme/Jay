package qsql

import (
	"qae/helper"
	"qae/libs/conf"
	"qae/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	MySQLdb *sqlx.DB
)

func InitMySQL() {
	var err error

	cfg := conf.GetMySQLConfig()
	MySQLdb, err = sqlx.Open("mysql", cfg.DSN())
	helper.EnsureOK(err, map[string]interface{}{
		"msg": "init mysql failed",
		"dsn": cfg.DSN(),
	})
	logger.Infof("Init MySQL Success")
}

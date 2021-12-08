package initial

import (
	"qae/env"
	"qae/libs/conf"
	"qae/libs/qkv"
	"qae/libs/qsql"
	"qae/logger"
)

func InitQAE() {
	env.InitEnv()
	logger.InitLogger()
	conf.InitStoreConfig()
	qsql.InitMySQL()
	qkv.InitLevelDB()
}

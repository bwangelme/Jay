package conf

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"qae/env"
	"qae/helper"
	"qae/logger"

	"gopkg.in/yaml.v3"
)

type MySQLConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

//DSN return dataSourceName
func (c *MySQLConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
}

type storeConfig struct {
	MySQL *MySQLConfig `yaml:"mysql"`
}

var (
	storeConf *storeConfig
)

//InitStoreConfig 初始化存储配置
func InitStoreConfig() {
	storeConf = &storeConfig{}

	confFile := filepath.Join(env.AppRoot, "config", "store.yaml")
	content, err := ioutil.ReadFile(confFile)
	helper.EnsureOK(err, map[string]interface{}{
		"msg":     "read conf file failed",
		"confile": confFile,
	})
	err = yaml.Unmarshal(content, storeConf)
	helper.EnsureOK(err, map[string]interface{}{
		"msg":     "parse config file failed",
		"content": content,
	})
	logger.Infof("Init Store Config Success")
}

func GetMySQLConfig() *MySQLConfig {
	return storeConf.MySQL
}

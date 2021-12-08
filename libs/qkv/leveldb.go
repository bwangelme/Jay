package qkv

import (
	"path/filepath"
	"qae/env"
	"qae/helper"
	"qae/logger"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	Bridge *leveldb.DB

	qaeFile = filepath.Join(env.AppRoot, "permdir", "qae.db")
)

func InitLevelDB() {
	var err error
	Bridge, err = leveldb.OpenFile(qaeFile, nil)
	helper.EnsureOK(err, map[string]interface{}{
		"msg":      "init leveldb failed",
		"filepath": qaeFile,
	})
	logger.Infof("Init LevelDB Success")
}

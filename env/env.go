package env

import (
	"log"
	"os"
)

var (
	//AppRoot app 的运行目录
	AppRoot string
)

const (
	keyAppRoot = "QAE_APPROOT"
)

func InitEnv() {
	AppRoot = ensureEnv(keyAppRoot)
	log.Println("Init Env Success")
}

func ensureEnv(envKey string) string {
	val := os.Getenv(envKey)
	if val == "" {
		log.Panicf("Get env %s failed", envKey)
	}

	return val
}

package appmodel

import (
	"time"
)

type App struct {
	id           int
	name         string
	owner        string
	repo         string
	createdTime  time.Time
	modifiedTime time.Time
	status       string
}

func AddApp(name string, owner string, repo string) {
}

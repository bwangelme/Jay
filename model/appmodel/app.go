package appmodel

import (
	"context"
	"qae/libs/qsql"
	"time"

	xerrors "github.com/pkg/errors"
)

type AppStatus int

const (
	Active   AppStatus = 1
	Deactive AppStatus = 2
)

type App struct {
	Id           int
	Name         string
	Owner        string
	Repo         string
	CreatedTime  time.Time
	ModifiedTime time.Time
	Status       int
}

func AddApp(ctx context.Context, name string, owner string, repo string, status AppStatus) (*App, error) {
	sql := `insert into apps(name, owner, repo, status) values (?, ?, ?, ?)`
	args := []interface{}{name, owner, repo, status}
	result, err := qsql.MySQLdb.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, xerrors.WithMessagef(err, "sql %v, args %v", sql, args)
	}

	appId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetApp(ctx, appId)
}

func GetApp(ctx context.Context, id int64) (*App, error) {
	var app App

	sql := `select id, name, owner, repo, created_time, modified_time, status from apps where id = ?`
	args := []interface{}{id}
	rows := qsql.MySQLdb.QueryRowContext(ctx, sql, args...)
	err := rows.Scan(&app.Id, &app.Name, &app.Owner, &app.Repo, &app.CreatedTime, &app.ModifiedTime, &app.Status)
	if err != nil {
		return nil, xerrors.WithMessagef(err, "sql %v, args %v", sql, args)
	}

	return &app, nil
}

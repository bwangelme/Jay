package appmodel

import (
	"context"
	"qae/libs/qsql"
	"time"

	"github.com/jmoiron/sqlx"
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
	CreatedTime  time.Time `db:"created_time"`
	ModifiedTime time.Time `db:"modified_time"`
	Status       int
}

func Add(ctx context.Context, name string, owner string, repo string, status AppStatus) (*App, error) {
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

	return Get(ctx, appId)
}

func List(ctx context.Context, start, limit int64) (int64, []int64, error) {
	var ids = make([]int64, 0)

	sql := `select id from apps limit ?,?`
	args := []interface{}{start, limit}
	err := qsql.MySQLdb.SelectContext(ctx, &ids, sql, args...)
	if err != nil {
		return 0, nil, xerrors.WithMessagef(err, "sql %v, args %v", sql, args)
	}

	var total int64
	err = qsql.MySQLdb.Get(&total, "SELECT count(*) from apps")
	if err != nil {
		return 0, nil, xerrors.WithMessagef(err, "sql %v, args %v", sql, args)
	}

	return total, ids, nil
}

func Get(ctx context.Context, id int64) (*App, error) {
	var app App

	sql := `select * from apps where id = ?`
	args := []interface{}{id}
	err := qsql.MySQLdb.SelectContext(ctx, &app, sql, args...)
	if err != nil {
		return nil, xerrors.WithMessagef(err, "sql %v, args %v", sql, args)
	}

	return &app, nil
}

func GetMulti(ctx context.Context, ids []int64) ([]*App, error) {
	var apps []App

	query, args, err := sqlx.In("select * from apps where id in (?);", ids)
	if err != nil {
		return nil, xerrors.WithMessagef(err, "args %v", args)
	}

	query = qsql.MySQLdb.Rebind(query)
	err = qsql.MySQLdb.SelectContext(ctx, &apps, query, args...)
	if err != nil {
		return nil, xerrors.WithMessagef(err, "sql: %v args %v", query, args)
	}

	var result = make([]*App, 0)
	for i := 0; i < len(apps); i++ {
		result = append(result, &apps[i])
	}

	return result, nil
}

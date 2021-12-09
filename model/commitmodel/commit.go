package commitmodel

import (
	"context"
	"fmt"
	"qae/libs/qkv"
	"qae/libs/qsql"
	"time"

	xerrors "github.com/pkg/errors"
)

type Commit struct {
	Id           int
	App          string
	GitCommit    string
	CreatedTime  time.Time
	ModifiedTime time.Time
	AppYAML      string
}

func genCommitKey(id int64) []byte {
	key := fmt.Sprintf("bridge/commit/%d", id)
	return []byte(key)
}

func AddCommit(ctx context.Context, app string, gitCommit string, appYAML string) (*Commit, error) {
	sql := `insert into commit(app, git_commit) values (?, ?)`
	args := []interface{}{app, gitCommit}

	result, err := qsql.MySQLdb.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, xerrors.WithMessagef(err, "sql %v, args %v", sql, args)
	}

	commitId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	key := genCommitKey(commitId)
	err = qkv.Bridge.Put(key, []byte(appYAML), nil)
	if err != nil {
		return nil, xerrors.WithMessagef(err, "key: %v", string(key))
	}

	return GetCommit(ctx, commitId)
}

func GetCommit(ctx context.Context, id int64) (*Commit, error) {
	var commit Commit

	sql := `select id, app, git_commit, created_time, modified_time from commit where id = ?`
	args := []interface{}{id}
	rows := qsql.MySQLdb.QueryRowContext(ctx, sql, args...)
	err := rows.Scan(&commit.Id, &commit.GitCommit, &commit.CreatedTime, &commit.CreatedTime, &commit.ModifiedTime)
	if err != nil {
		return nil, xerrors.WithMessagef(err, "sql %v, args %v", sql, args)
	}

	key := genCommitKey(id)
	appYAML, err := qkv.Bridge.Get(key, nil)
	if err != nil {
		return nil, xerrors.WithMessagef(err, "key %v", string(key))
	}

	commit.AppYAML = string(appYAML)
	return &commit, nil
}

package grpcview

import (
	"context"
	"qae/pb/gen/commit"
)

type CommitHandler struct {
	commit.UnimplementedCommitServiceServer
}

var (
	Commit *CommitHandler
)

func InitCommitHandler() {
	Commit = &CommitHandler{}
}

//Get 获取 Commit
func (s *CommitHandler) Get(ctx context.Context, req *commit.CommitID) (resp *commit.Commit, err error) {
	return nil, nil
}

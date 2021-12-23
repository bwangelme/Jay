package grpcview

import (
	"context"
	pbc "qae/pb/gen/commit"
	"qae/services/commitservice"
)

type CommitHandler struct {
	pbc.UnimplementedCommitServiceServer
}

var (
	Commit *CommitHandler
)

func InitCommitHandler() {
	Commit = &CommitHandler{}
}

//Get 获取 Commit
func (s *CommitHandler) Get(ctx context.Context, req *pbc.CommitID) (resp *pbc.Commit, err error) {
	return nil, nil
}

func (s *CommitHandler) Create(ctx context.Context, req *pbc.CreateReq) (*pbc.CreateResp, error) {
	commit, err := commitservice.Add(req.App, req.GitCommit, req.AppYAML)
	return commit, err
}

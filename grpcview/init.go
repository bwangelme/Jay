package grpcview

import (
	"qae/pb/gen/app"
	"qae/pb/gen/commit"

	"google.golang.org/grpc"
)

func InitHandlers(server *grpc.Server) {
	initAppHandler()
	InitCommitHandler()

	app.RegisterAppServiceServer(server, App)
	commit.RegisterCommitServiceServer(server, Commit)
}

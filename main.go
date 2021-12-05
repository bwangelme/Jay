package main

import (
	"net"
	"qae/grpcview"
	"qae/initial"
	"qae/logger"
	"qae/pb/gen/app"

	"google.golang.org/grpc"
)

func main() {
	initial.InitQAE()
	addr := ":8000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Infof("Listen %v failed %v\n", addr, err)
		return
	}

	s := grpc.NewServer()

	// register view
	grpcview.InitHandlers()
	app.RegisterAppServiceServer(s, grpcview.App)

	if err := s.Serve(listener); err != nil {
		logger.Infof("grpc serve failed %v", err)
		return
	}
}

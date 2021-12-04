package main

import (
	"context"
	"net"
	"qae/initial"
	"qae/logger"
	"qae/pb/gen/app"

	"google.golang.org/grpc"
)

type server struct {
	app.UnimplementedAppServiceServer
}

func NewServer() *server {
	return &server{}
}

//Get 获取 App
func (s *server) Get(ctx context.Context, req *app.AppID) (resp *app.App, err error) {
	return &app.App{
		Id:    "1",
		Name:  "bridge",
		Owner: "xuyundong",
	}, nil
}

func (s *server) Triple(ctx context.Context, req *app.Number) (resp *app.Number, err error) {
	return &app.Number{
		Value: req.Value * 3,
	}, nil
}

func main() {
	initial.InitQAE()
	addr := ":8000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Infof("Listen %v failed %v\n", addr, err)
		return
	}

	s := grpc.NewServer()
	app.RegisterAppServiceServer(s, NewServer())
	if err := s.Serve(listener); err != nil {
		logger.Infof("grpc serve failed %v", err)
		return
	}
}

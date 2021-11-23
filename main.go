package main

import (
	"context"
	"log"
	"net"
	"qae/proto/app"

	"google.golang.org/grpc"
)

type server struct {
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

func main() {
	addr := ":8000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("Listen %v failed %v\n", addr, err)
		return
	}

	s := grpc.NewServer()
}

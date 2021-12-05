package grpcview

import (
	"context"
	"fmt"
	"qae/grpcview/convertor"
	"qae/grpcview/validator"
	"qae/pb/gen/app"
	"qae/services/appservice"
)

var (
	App *AppHandler
)

type AppHandler struct {
	app.UnimplementedAppServiceServer
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

//Get 获取 App
func (s *AppHandler) Get(ctx context.Context, req *app.AppID) (resp *app.App, err error) {
	return &app.App{
		Id:    "1",
		Name:  "bridge",
		Owner: "xuyundong",
	}, nil
}

func (s *AppHandler) Create(ctx context.Context, req *app.CreateReq) (resp *app.CreateResp, err error) {
	resp = &app.CreateResp{}
	err = validator.ValidateAppName(req.Name)

	if err != nil {
		return returnGRPCError(err), nil
	}
	err = validator.ValidateAppRepo(req.Repo)
	if err != nil {
		return returnGRPCError(err), nil
	}

	res, err := appservice.CreateApp(ctx, req.Name, req.Repo)
	if err != nil {
		return returnGRPCError(err), nil
	}

	resp.App = convertor.NewGRPCAppFromModelApp(res)
	resp.Err = ""
	return resp, nil
}

func (s *AppHandler) Triple(ctx context.Context, req *app.Number) (resp *app.Number, err error) {
	return &app.Number{
		Value: req.Value * 3,
	}, nil
}

func initAppHandler() {
	App = NewAppHandler()
}

func InitHandlers() {
	initAppHandler()
}

func returnGRPCError(err error) (resp *app.CreateResp) {
	resp = &app.CreateResp{}
	resp.Err = fmt.Sprintf("%v", err)
	resp.App = nil
	return resp
}

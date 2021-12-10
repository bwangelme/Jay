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

func (s *AppHandler) List(ctx context.Context, req *app.ListReq) (resp *app.ListResp, err error) {
	resp = &app.ListResp{}

	total, apps, err := appservice.List(ctx, req.Start, req.Limit)

	var grpcApps = make([]*app.App, 0)
	for i := 0; i < len(apps); i++ {
		a := convertor.NewGRPCAppFromModelApp(apps[i])
		grpcApps = append(grpcApps, a)
	}
	resp.Total = total
	resp.Apps = grpcApps

	return resp, nil
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

	res, err := appservice.Create(ctx, req.Name, req.Repo)
	if err != nil {
		return returnGRPCError(err), nil
	}

	resp.App = convertor.NewGRPCAppFromModelApp(res)
	resp.Err = ""
	return resp, nil
}

func (s *AppHandler) Delete(ctx context.Context, req *app.DeleteReq) (*app.DeleteResp, error) {
	var resp = &app.DeleteResp{}
	cnt, err := appservice.Delete(ctx, req.Ids)
	if err != nil {
		resp.Err = err.Error()
		return resp, nil
	}

	resp.Err = ""
	resp.Cnt = cnt
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

func returnGRPCError(err error) (resp *app.CreateResp) {
	resp = &app.CreateResp{}
	resp.Err = fmt.Sprintf("%v", err)
	resp.App = nil
	return resp
}

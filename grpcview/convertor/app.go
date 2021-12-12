package convertor

import (
	"qae/model/appmodel"
	"qae/pb/gen/app"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewGRPCAppFromModelApp(res *appmodel.App) *app.App {
	return &app.App{
		Id:          res.Id,
		Name:        res.Name,
		Owner:       res.Owner,
		Repo:        res.Repo,
		CreatedTime: timestamppb.New(res.CreatedTime),
	}
}

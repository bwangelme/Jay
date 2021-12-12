package appservice

import (
	"context"
	"qae/model/appmodel"
)

const OWNER = "bwangel"

func Create(ctx context.Context, name string, repo string) (*appmodel.App, error) {
	return appmodel.Add(ctx, name, OWNER, repo, appmodel.Active)
}

func Delete(ctx context.Context, ids []int64) (int64, error) {
	return appmodel.Delete(ctx, ids)
}

func List(ctx context.Context, start, limit int64) (total int64, apps []*appmodel.App, err error) {
	total, ids, err := appmodel.List(ctx, start, limit)
	if err != nil {
		return 0, nil, err
	}

	apps, err = appmodel.GetMulti(ctx, ids)
	if err != nil {
		return 0, nil, err
	}

	return total, apps, nil
}

func Get(ctx context.Context, id int64) (*appmodel.App, error) {
	return appmodel.Get(ctx, id)
}

func GetMulti(ctx context.Context, ids []int64) ([]*appmodel.App, error) {
	return appmodel.GetMulti(ctx, ids)
}

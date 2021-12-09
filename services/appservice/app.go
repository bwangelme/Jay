package appservice

import (
	"context"
	"qae/model/appmodel"
)

const OWNER = "bwangel"

func Create(ctx context.Context, name string, repo string) (*appmodel.App, error) {
	return appmodel.Add(ctx, name, OWNER, repo, appmodel.Active)
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

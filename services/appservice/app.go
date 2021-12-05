package appservice

import (
	"context"
	"qae/model/appmodel"
)

const OWNER = "bwangel"

func CreateApp(ctx context.Context, name string, repo string) (*appmodel.App, error) {
	return appmodel.AddApp(ctx, name, OWNER, repo, appmodel.Active)
}


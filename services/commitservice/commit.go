package commitservice

import (
	"context"
	"qae/model/commitmodel"
)

func Add(ctx context.Context, app string, gitCommit string, appYAML string) (*commitmodel.Commit, error) {
	// TODO: 在 service 层将 gitcommit 转换成 giturl
	return commitmodel.Add(ctx, app, gitCommit, appYAML)
}

package infra

import "context"

func New() Container {
	return Container{
		Fs:  &FsRepository{},
		Log: &LogRepository{},
		Cmd: &CmdRepository{},
	}
}

type Container struct {
	Fs  FsRepositoryInterface
	Log LogRepositoryInterface
	Cmd CmdRepositoryInterface
}

type reposKey struct{}

func NewContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, reposKey{}, New())

	return ctx
}

// Deprecated
func Use(ctx context.Context) Container {
	repos, ok := ctx.Value(reposKey{}).(Container)
	if !ok {
		return New()
	}
	return repos
}

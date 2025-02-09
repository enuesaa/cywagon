package repository

import "context"

func NewRepos() Repos {
	return Repos{
		Fs:  &FsRepository{},
		Log: &LogRepository{},
		Cmd: &CmdRepository{},
	}
}

type Repos struct {
	Fs  FsRepositoryInterface
	Log LogRepositoryInterface
	Cmd CmdRepositoryInterface
}

type reposKey struct{}

func NewContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, reposKey{}, NewRepos())

	return ctx
}

func Use(ctx context.Context) Repos {
	repos, ok := ctx.Value(reposKey{}).(Repos)
	if !ok {
		return NewRepos()
	}
	return repos
}

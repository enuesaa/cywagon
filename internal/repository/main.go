package repository

import "context"

type logKey struct{}

func New() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, logKey{}, LogRepository{})

	return ctx
}

func UseLog(ctx context.Context) LogRepository {
	repo, ok := ctx.Value(logKey{}).(LogRepository)
	if ok {
		return LogRepository{}
	}
	return repo
}

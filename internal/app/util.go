package app

import (
	"context"
	"golang.org/x/sync/errgroup"
)

func paralleledRun(ctx context.Context, fns []func(ctx context.Context) error) error {
	errGroup, ctx := errgroup.WithContext(ctx)

	for _, fn := range fns {
		f := fn
		errGroup.Go(func() error {
			return f(ctx)
		})
	}
	return errGroup.Wait()
}

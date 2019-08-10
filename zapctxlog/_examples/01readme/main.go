package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/podhmo/ctxlog"
	"github.com/podhmo/ctxlog/zapctxlog"
)

func main() {
	l := zapctxlog.MustNew()
	ctx, _ := ctxlog.WithLogger(context.Background(), l).With("x-id", 10)
	if err := f(ctx); err != nil {
		ctxlog.Get(ctx).WithError(err).Warning("!!")
	}
}

func f(ctx context.Context) error {
	return errors.Wrap(fmt.Errorf("hmm"), "on f")
}

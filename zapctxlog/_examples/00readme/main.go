package main

import (
	"context"
	"log"

	"github.com/podhmo/ctxlog"
	"github.com/podhmo/ctxlog/zapctxlog"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	l := zapctxlog.MustNew()
	ctx, _ := ctxlog.WithLogger(context.Background(), l).With("x-id", 10)
	return f(ctx)
}

func f(ctx context.Context) error {
	ctx, log := ctxlog.Get(ctx).With("y-id", 20)
	log.Info("start f")
	defer log.Info("end f")
	return g(ctx)
}

func g(ctx context.Context) error {
	ctx, log := ctxlog.Get(ctx).With("y-id", 20)
	log.Info("start g")
	defer log.Info("end g")
	return nil
}

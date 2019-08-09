package ctxlog_test

import (
	"context"

	"github.com/podhmo/ctxlog"
	"github.com/podhmo/ctxlog/zapctxlog"
	"go.uber.org/zap"
)

func Example() {
	var f, g func(context.Context) error
	f = func(ctx context.Context) error {
		ctx, log := ctxlog.Logger(ctx).With("x-id", 10)
		log.Info("start f")
		defer log.Info("end f")
		return g(ctx)
	}
	g = func(ctx context.Context) error {
		_, log := ctxlog.Logger(ctx).With("y-id", 20)
		log.Info("start g")
		defer log.Info("end g")
		return nil
	}

	log, _ := zapctxlog.New(zapctxlog.WithNewInternal(newZapLogger))
	_ = f(ctxlog.WithLogger(context.Background(), log))

	// Output:
	// INFO	ctxlog/example_test.go:15	start f	{"x-id": 10}
	// INFO	ctxlog/example_test.go:21	start g	{"x-id": 10, "y-id": 20}
	// INFO	ctxlog/example_test.go:23	end g	{"x-id": 10, "y-id": 20}
	// INFO	ctxlog/example_test.go:17	end f	{"x-id": 10}
}

func newZapLogger(options ...zap.Option) *zap.Logger {
	c := zap.NewDevelopmentConfig()
	c.EncoderConfig.TimeKey = ""
	c.OutputPaths = []string{"stdout"}
	c.ErrorOutputPaths = []string{"stdout"}
	l, _ := c.Build(options...)
	return l
}

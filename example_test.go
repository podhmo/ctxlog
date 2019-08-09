package ctxlog

import (
	"context"

	"github.com/podhmo/ctxlog/zapctxlog"
	"go.uber.org/zap"
)

func Example() {
	var f, g func(context.Context) error
	f = func(ctx context.Context) error {
		ctx, log := Get(ctx).With("x-id", 10)
		log.Info("start f")
		defer log.Info("end f")
		return g(ctx)
	}
	g = func(ctx context.Context) error {
		ctx, log := Get(ctx).With("y-id", 20)
		log.Info("start g")
		defer log.Info("end g")
		return nil
	}

	log, _ := zapctxlog.New(zapctxlog.WithNewInternal(newZapLogger))
	f(WithLogger(context.Background(), log))

	// Output:
	// INFO	ctxlog/example_test.go:14	start f	{"x-id": 10}
	// INFO	ctxlog/example_test.go:20	start g	{"x-id": 10, "y-id": 20}
	// INFO	ctxlog/example_test.go:22	end g	{"x-id": 10, "y-id": 20}
	// INFO	ctxlog/example_test.go:16	end f	{"x-id": 10}
}

func newZapLogger(options ...zap.Option) *zap.Logger {
	c := zap.NewDevelopmentConfig()
	c.EncoderConfig.TimeKey = ""
	c.OutputPaths = []string{"stdout"}
	c.ErrorOutputPaths = []string{"stdout"}
	l, _ := c.Build(options...)
	return l
}

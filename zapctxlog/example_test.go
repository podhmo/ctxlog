package zapctxlog_test

import (
	"context"
	"time"

	"github.com/podhmo/ctxlog"
	"github.com/podhmo/ctxlog/zapctxlog"
	"go.uber.org/zap"
)

func Example() {
	var f, g func(context.Context) error
	f = func(ctx context.Context) error {
		ctx, log := ctxlog.Get(ctx).With("x-id", 10)
		log.Info("start f")
		defer log.Info("end f")
		return g(ctx)
	}
	g = func(ctx context.Context) error {
		_, log := ctxlog.Get(ctx).With("y-id", 20)
		log.Info("start g")
		defer log.Info("end g")
		return nil
	}

	log, _ := zapctxlog.New(zapctxlog.WithNewInternal(newZapLogger))
	_ = f(ctxlog.WithLogger(context.Background(), log))

	// Output:
	// INFO	zapctxlog/example_test.go:16	start f	{"x-id": 10}
	// INFO	zapctxlog/example_test.go:22	start g	{"x-id": 10, "y-id": 20}
	// INFO	zapctxlog/example_test.go:24	end g	{"x-id": 10, "y-id": 20}
	// INFO	zapctxlog/example_test.go:18	end f	{"x-id": 10}
}

func ExampleNew() {
	log, _ := zapctxlog.New(zapctxlog.WithNewInternal(newZapLogger))
	now, _ := time.Parse(time.RFC3339, "2011-11-11T22:22:22Z")
	log.With("now", now).Info("hello")

	// Output:
	// INFO	zapctxlog/example_test.go:40	hello	{"now": "2011-11-11T22:22:22.000Z"}
}

func newZapLogger(options ...zap.Option) *zap.Logger {
	c := zap.NewDevelopmentConfig()
	c.EncoderConfig.TimeKey = ""
	c.OutputPaths = []string{"stdout"}
	c.ErrorOutputPaths = []string{"stdout"}
	l, _ := c.Build(options...)
	return l
}

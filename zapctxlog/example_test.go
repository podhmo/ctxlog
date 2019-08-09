package zapctxlog

import (
	"time"

	"go.uber.org/zap"
)

func Example() {
	log, _, teardown := New(
		WithNewInternal(NewExample),
	)

	defer teardown()
	now, _ := time.Parse(time.RFC3339, "2011-11-11T22:22:22Z")
	log.With("now", now).Info("hello")

	// Output:
	// INFO    zapctxlog/example_test.go:17    hello     {"now": "2011-11-11T22:22:22.000Z"}
}

func NewExample(options ...zap.Option) *zap.Logger {
	c := zap.NewDevelopmentConfig()
	c.EncoderConfig.TimeKey = ""
	c.OutputPaths = []string{"stdout"}
	c.ErrorOutputPaths = []string{"stdout"}
	l, _ := c.Build(options...)
	return l
}

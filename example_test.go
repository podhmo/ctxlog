package ctxlog_test

import (
	"time"

	"github.com/podhmo/ctxlog"
)

func ExampleNoopLogger() {
	log := &ctxlog.NoopLogger{}

	now, _ := time.Parse(time.RFC3339, "2011-11-11T22:22:22Z")
	log.With("now", now).Info("hello")

	// Output:
	//
}

package stdctxlog_test

import (
	"log"
	"time"

	"github.com/podhmo/ctxlog/stdctxlog"
)

func Example() {
	log := stdctxlog.New(
		stdctxlog.WithFlags(log.Lshortfile),
	)
	now, _ := time.Parse(time.RFC3339, "2011-11-11T22:22:22Z")
	log.With("now", now).Info("hello")

	// Output:
	// example_test.go:15: hello	now:2011-11-11 22:22:22 +0000 UTC
}

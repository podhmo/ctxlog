package noopctxlog

import (
	"time"
)

func Example() {
	log := &Logger{}

	now, _ := time.Parse(time.RFC3339, "2011-11-11T22:22:22Z")
	log.With("now", now).Info("hello")

	// Output:
	//
}

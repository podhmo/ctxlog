package ctxlog

import (
	"fmt"
	"os"
	"sync"

	"github.com/podhmo/ctxlog/ctxlogcore"
)

var (
	once sync.Once
)

// getNoop :
func getNoop() ctxlogcore.Logger {
	once.Do(func() {
		if os.Getenv("CTXLOG_STRICT") != "" {
			panic("logger not set")
		}
		fmt.Fprintln(os.Stderr, "\x1b[33m**CTXLOG WARNING*************************")
		fmt.Fprintln(os.Stderr, "ctxlog.Logger is not found. please set logger, via ctxlog.WithLogger()")
		fmt.Fprintln(os.Stderr, "****************************************\x1b[0m")
	})
	return &NoopLogger{}
}

// NoopLogger :
type NoopLogger struct {
	KeysAndValues []interface{}
}

// With :
func (l *NoopLogger) With(keysAndValues ...interface{}) ctxlogcore.Logger {
	return &NoopLogger{
		KeysAndValues: append(l.KeysAndValues, keysAndValues...),
	}
}

// Debug :
func (l *NoopLogger) Debug(msg string) {
}

// Info :
func (l *NoopLogger) Info(msg string) {
}

// Warning :
func (l *NoopLogger) Warning(msg string) {
}

// Error :
func (l *NoopLogger) Error(msg string) {
}

// Fatal :
func (l *NoopLogger) Fatal(msg string) {
}

// Panic :
func (l *NoopLogger) Panic(msg string) {
}

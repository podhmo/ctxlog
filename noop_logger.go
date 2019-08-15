package ctxlog

import (
	"fmt"
	"os"
	"sync"
)

var (
	once sync.Once
)

// getNoop :
func getNoop() Logger {
	once.Do(func() {
		if os.Getenv("CTXLOG_STRICT") != "" {
			panic("logger not set")
		}
		fmt.Fprintln(os.Stderr, "\x1b[33m**CTXLOG WARNING*************************")
		fmt.Fprintln(os.Stderr, "ctxlog.Get is not found. please set logger, via ctxlog.Set()")
		fmt.Fprintln(os.Stderr, "****************************************\x1b[0m")
	})
	return &NoopLogger{}
}

// NoopLogger :
type NoopLogger struct {
	KeysAndValues []interface{}
}

// With :
func (l *NoopLogger) With(keysAndValues ...interface{}) Logger {
	return &NoopLogger{
		KeysAndValues: append(l.KeysAndValues, keysAndValues...),
	}
}

// WithError :
func (l *NoopLogger) WithError(err error) Logger {
	return l.With("error", err)
}

// Debug :
func (l *NoopLogger) Debug(msg string, keysAndValues ...interface{}) {
}

// Info :
func (l *NoopLogger) Info(msg string, keysAndValues ...interface{}) {
}

// Warning :
func (l *NoopLogger) Warning(msg string, keysAndValues ...interface{}) {
}

// Error :
func (l *NoopLogger) Error(msg string, keysAndValues ...interface{}) {
}

// Fatal :
func (l *NoopLogger) Fatal(msg string, keysAndValues ...interface{}) {
}

// Panic :
func (l *NoopLogger) Panic(msg string, keysAndValues ...interface{}) {
}

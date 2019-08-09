package ctxlog

import (
	"os"
	"sync"
)

var (
	once sync.Once
)

// getNoop :
func getNoop() Logger {
	once.Do(func() {
		if os.Getenv("CTXLOG_LOOSE") == "" {
			panic("logger not set")
		}
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

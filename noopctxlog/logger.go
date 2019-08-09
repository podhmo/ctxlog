package noopctxlog

import (
	"os"
	"sync"

	"github.com/podhmo/ctxlog/ctxlogcore"
)

var (
	once sync.Once
)

// Get :
func Get() ctxlogcore.Logger {
	once.Do(func() {
		if os.Getenv("CTXLOG_LOOSE") == "" {
			panic("logger not set")
		}
	})
	return &Logger{}
}

// Logger :
type Logger struct {
	KeysAndValues []interface{}
}

// With :
func (l *Logger) With(keysAndValues ...interface{}) ctxlogcore.Logger {
	return &Logger{
		KeysAndValues: append(l.KeysAndValues, keysAndValues...),
	}
}

// Debug :
func (l *Logger) Debug(msg string) {
}

// Info :
func (l *Logger) Info(msg string) {
}

// Warning :
func (l *Logger) Warning(msg string) {
}

// Error :
func (l *Logger) Error(msg string) {
}

// Fatal :
func (l *Logger) Fatal(msg string) {
}

// Panic :
func (l *Logger) Panic(msg string) {
}

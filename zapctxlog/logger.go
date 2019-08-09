package zapctxlog

import (
	"github.com/podhmo/ctxlog"
	"go.uber.org/zap"
)

// New :
func New() (ctxlog.Logger, error, func() error) {
	logger, err := zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err, nil
	}
	sugar := logger.Sugar()
	return &Logger{
		Internal: sugar,
	}, nil, logger.Sync
}

// Logger :
type Logger struct {
	Internal      *zap.SugaredLogger
	KeysAndValues []interface{}
}

// With :
func (l *Logger) With(k string, v interface{}) ctxlog.Logger {
	return &Logger{
		Internal:      l.Internal,
		KeysAndValues: append(l.KeysAndValues, k, v),
	}
}

// Debug :
func (l *Logger) Debug(msg string) {
	l.Internal.Debugw(msg, l.KeysAndValues...)
}

// Info :
func (l *Logger) Info(msg string) {
	l.Internal.Infow(msg, l.KeysAndValues...)
}

// Warning :
func (l *Logger) Warning(msg string) {
	l.Internal.Warnw(msg, l.KeysAndValues...)
}

// Error :
func (l *Logger) Error(msg string) {
	l.Internal.Errorw(msg, l.KeysAndValues...)
}

// Fatal :
func (l *Logger) Fatal(msg string) {
	l.Internal.Fatalw(msg, l.KeysAndValues...)
}

// Panic :
func (l *Logger) Panic(msg string) {
	l.Internal.Panicw(msg, l.KeysAndValues...)
}

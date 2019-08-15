package zapctxlog

import (
	"github.com/podhmo/ctxlog"
	"go.uber.org/zap"
)

// Config :
type Config struct {
	NewInternal func(...zap.Option) (*zap.Logger, error)
	Options     []zap.Option
}

// WithNewInternal :
func WithNewInternal(newInternal func(...zap.Option) *zap.Logger) func(*Config) {
	return func(c *Config) {
		c.NewInternal = func(options ...zap.Option) (*zap.Logger, error) {
			l := newInternal(options...)
			return l, nil
		}
	}
}

// WithNewInternal2 :
func WithNewInternal2(newInternal func(...zap.Option) (*zap.Logger, error)) func(*Config) {
	return func(c *Config) {
		c.NewInternal = newInternal
	}
}

// WithOption :
func WithOption(options ...zap.Option) func(*Config) {
	return func(c *Config) {
		c.Options = append(c.Options, options...)
	}
}

// MustNew :
func MustNew(options ...func(*Config)) ctxlog.Logger {
	l, err := New(options...)
	if err != nil {
		panic(err)
	}
	return l
}

// New :
func New(options ...func(*Config)) (ctxlog.Logger, error) {
	c := &Config{
		NewInternal: zap.NewProduction,
		Options:     []zap.Option{zap.AddCallerSkip(1)},
	}
	for _, opt := range options {
		opt(c)
	}
	logger, err := c.NewInternal(c.Options...)
	if err != nil {
		return nil, err
	}
	sugar := logger.Sugar()
	return &Logger{
		Internal: sugar,
	}, nil
}

// Logger :
type Logger struct {
	Internal      *zap.SugaredLogger
	KeysAndValues []interface{}
}

// With :
func (l *Logger) With(keysAndValues ...interface{}) ctxlog.Logger {
	return &Logger{
		Internal:      l.Internal,
		KeysAndValues: append(l.KeysAndValues, keysAndValues...),
	}
}

// WithError :
func (l *Logger) WithError(err error) ctxlog.Logger {
	return l.With("error", err)
}

// Debug :
func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	l.Internal.Debugw(msg, append(l.KeysAndValues, keysAndValues...)...)
}

// Info :
func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.Internal.Infow(msg, append(l.KeysAndValues, keysAndValues...)...)
}

// Warning :
func (l *Logger) Warning(msg string, keysAndValues ...interface{}) {
	l.Internal.Warnw(msg, append(l.KeysAndValues, keysAndValues...)...)
}

// Error :
func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	l.Internal.Errorw(msg, append(l.KeysAndValues, keysAndValues...)...)
}

// Fatal :
func (l *Logger) Fatal(msg string, keysAndValues ...interface{}) {
	l.Internal.Fatalw(msg, append(l.KeysAndValues, keysAndValues...)...)
}

// Panic :
func (l *Logger) Panic(msg string, keysAndValues ...interface{}) {
	l.Internal.Panicw(msg, append(l.KeysAndValues, keysAndValues...)...)
}

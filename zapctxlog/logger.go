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

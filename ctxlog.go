package ctxlog

import (
	"context"
	"fmt"

	"github.com/podhmo/ctxlog/ctxlogcore"
)

type ctxKeyType string

const (
	ctxKey = ctxKeyType("ctxlog")
)

// Logger :
type Logger = ctxlogcore.Logger

// Get :
func Get(ctx context.Context) *LoggerContext {
	if ctx, ok := ctx.(*LoggerContext); ok {
		return ctx
	}

	v := ctx.Value(ctxKey)
	l, ok := (v).(ctxlogcore.Logger)
	if !ok {
		l = getNoop()
	}
	return &LoggerContext{Context: ctx, Logger: l}
}

// Set
func Set(ctx context.Context, l ctxlogcore.Logger) *LoggerContext {
	return &LoggerContext{
		Context: context.WithValue(ctx, ctxKey, l),
		Logger:  l,
	}
}

// LoggerContext :
type LoggerContext struct {
	context.Context
	ctxlogcore.Logger
}

// With :
func (lc *LoggerContext) With(keysAndValues ...interface{}) (context.Context, *LoggerContext) {
	l := lc.Logger
	for i := 0; i < len(keysAndValues); i += 2 {
		k := keysAndValues[i]
		v := keysAndValues[i+1]
		switch k := k.(type) {
		case string:
			l = l.With(k, v)
		default:
			l = l.With(fmt.Sprintf("%s", k), v)
		}
	}
	ctx := Set(lc.Context, l)
	return ctx, ctx
}

// WithError :
func (lc *LoggerContext) WithError(err error) ctxlogcore.Logger {
	return lc.Logger.With("error", err)
}

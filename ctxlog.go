package ctxlog

import (
	"context"

	"github.com/podhmo/ctxlog/ctxlogcore"
	"github.com/podhmo/ctxlog/noopctxlog"
)

type ctxKeyType string

const (
	ctxKey = ctxKeyType("ctxlog")
)

// Get :
func Get(ctx context.Context) Logger {
	v := ctx.Value(ctxKey)
	if l, ok := (v).(Logger); ok {
		return l
	}
	return noopctxlog.Get()
}

// Logger :
type Logger = ctxlogcore.Logger

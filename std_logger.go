package ctxlog

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/podhmo/ctxlog/ctxlogcore"
)

// FIXME: bad implementation

// StdLogger :
type StdLogger struct {
	w             io.Writer
	Internal      *log.Logger
	KeysAndValues []interface{}
}

// NewStdLogger :
func NewStdLogger() *StdLogger {
	var logger *StdLogger
	w := os.Stdout
	output := &LTSVOutput{W: w, KeyAndValues: func() []interface{} { return logger.KeysAndValues }}
	logger = &StdLogger{w: w, Internal: log.New(output, "", log.LstdFlags|log.Lshortfile)}
	return logger
}

// With :
func (l *StdLogger) With(keysAndValues ...interface{}) ctxlogcore.Logger {
	w := l.w
	var logger *StdLogger
	output := &LTSVOutput{W: w, KeyAndValues: func() []interface{} { return logger.KeysAndValues }}
	logger = &StdLogger{
		w:             w,
		Internal:      log.New(output, l.Internal.Prefix(), l.Internal.Flags()),
		KeysAndValues: append(l.KeysAndValues, keysAndValues...),
	}
	return logger
}

// Debug :
func (l *StdLogger) Debug(msg string) {
	if err := l.Internal.Output(2, msg); err != nil {
		panic(err)
	}
}

// Info :
func (l *StdLogger) Info(msg string) {
	if err := l.Internal.Output(2, msg); err != nil {
		panic(err)
	}
}

// Warning :
func (l *StdLogger) Warning(msg string) {
	if err := l.Internal.Output(2, msg); err != nil {
		panic(err)
	}
}

// Error :
func (l *StdLogger) Error(msg string) {
	if err := l.Internal.Output(2, msg); err != nil {
		panic(err)
	}
}

// Fatal :
func (l *StdLogger) Fatal(msg string) {
	if err := l.Internal.Output(2, msg); err != nil {
		panic(err)
	}
	os.Exit(1)
}

// Panic :
func (l *StdLogger) Panic(msg string) {
	if err := l.Internal.Output(2, msg); err != nil {
		panic(err)
	}
	panic(msg)
}

// LTSVOutput :
type LTSVOutput struct {
	W            io.Writer
	KeyAndValues func() []interface{}
}

// Write :
func (o *LTSVOutput) Write(p []byte) (n int, err error) {
	keyAndValues := o.KeyAndValues()
	if len(keyAndValues) == 0 {
		return o.W.Write(p)
	}
	if len(p) == 0 {
		return 0, nil
	}
	if p[len(p)-1] == '\n' {
		p = p[:len(p)-1]
	}
	n, err = o.W.Write(p)
	if err != nil {
		return n, err
	}

	for i := 0; i < len(keyAndValues); i += 2 {
		k := keyAndValues[i]
		v := keyAndValues[i+1]
		m, err := fmt.Printf("	%s:%v", k, v)
		n += m
		if err != nil {
			return n, err
		}
	}
	m, err := o.W.Write([]byte{'\n'})
	if err != nil {
		return n, err
	}
	n += m
	return n, nil
}

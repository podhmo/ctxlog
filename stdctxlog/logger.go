package stdctxlog

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/podhmo/ctxlog"
)

// FIXME: bad implementation

// Config :
type Config struct {
	Writer  io.Writer
	Flag    int
	Verbose bool
}

// WithFlags :
func WithFlags(flag int) func(*Config) {
	return func(c *Config) {
		c.Flag = flag
	}
}

// WithVerbose :
func WithVerbose() func(*Config) {
	return func(c *Config) {
		c.Verbose = true
	}
}

// Logger :
type Logger struct {
	Config        *Config
	Internal      *log.Logger
	KeysAndValues []interface{}
}

// New :
func New(options ...func(*Config)) *Logger {
	c := &Config{
		Writer: os.Stdout,
		Flag:   log.LstdFlags | log.Lshortfile,
	}
	for _, opt := range options {
		opt(c)
	}

	var logger *Logger
	output := &LTSVOutput{
		W:            c.Writer,
		KeyAndValues: func() []interface{} { return logger.KeysAndValues },
		Verbose:      c.Verbose,
	}
	logger = &Logger{
		Config:   c,
		Internal: log.New(output, "", c.Flag),
	}
	return logger
}

// With :
func (l *Logger) With(keysAndValues ...interface{}) ctxlog.Logger {
	var logger *Logger
	output := &LTSVOutput{
		W:            l.Config.Writer,
		KeyAndValues: func() []interface{} { return logger.KeysAndValues },
		Verbose:      l.Config.Verbose,
	}
	logger = &Logger{
		Config:        l.Config,
		Internal:      log.New(output, l.Internal.Prefix(), l.Internal.Flags()),
		KeysAndValues: append(l.KeysAndValues, keysAndValues...),
	}
	return logger
}

// WithError :
func (l *Logger) WithError(err error) ctxlog.Logger {
	return l.With("error", err)
}

// Debug :
func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	// TODO: support level
	wl := l
	if len(keysAndValues) > 0 {
		wl = l.With(keysAndValues...).(*Logger)
	}
	if err := wl.Internal.Output(2, msg); err != nil {
		panic(err)
	}
}

// Info :
func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	// TODO: support level
	wl := l
	if len(keysAndValues) > 0 {
		wl = l.With(keysAndValues...).(*Logger)
	}
	if err := wl.Internal.Output(2, msg); err != nil {
		panic(err)
	}
}

// Warning :
func (l *Logger) Warning(msg string, keysAndValues ...interface{}) {
	// TODO: support level
	wl := l
	if len(keysAndValues) > 0 {
		wl = l.With(keysAndValues...).(*Logger)
	}
	if err := wl.Internal.Output(2, msg); err != nil {
		panic(err)
	}
}

// Error :
func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	// TODO: support level
	wl := l
	if len(keysAndValues) > 0 {
		wl = l.With(keysAndValues...).(*Logger)
	}
	if err := wl.Internal.Output(2, msg); err != nil {
		panic(err)
	}
}

// Fatal :
func (l *Logger) Fatal(msg string, keysAndValues ...interface{}) {
	// TODO: support level
	wl := l
	if len(keysAndValues) > 0 {
		wl = l.With(keysAndValues...).(*Logger)
	}
	if err := wl.Internal.Output(2, msg); err != nil {
		panic(err)
	}
	os.Exit(1)
}

// Panic :
func (l *Logger) Panic(msg string, keysAndValues ...interface{}) {
	// TODO: support level
	wl := l
	if len(keysAndValues) > 0 {
		wl = l.With(keysAndValues...).(*Logger)
	}
	if err := wl.Internal.Output(2, msg); err != nil {
		panic(err)
	}
	panic(msg)
}

// LTSVOutput :
type LTSVOutput struct {
	W            io.Writer
	Verbose      bool
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

	msgfmt := "	%s:%v"
	if o.Verbose {
		msgfmt = "	%s:%+v"
	}

	for i := 0; i < len(keyAndValues); i += 2 {
		k := keyAndValues[i]
		v := keyAndValues[i+1]
		m, err := fmt.Printf(msgfmt, k, v)
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

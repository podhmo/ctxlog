package ctxlogcore

// Logger :
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Fatal(msg string)
	Panic(msg string) // panic

	// structual
	With(keysAndValues ...interface{}) Logger
}

# ctxlog

contextual logging and structual logging

main.go

```go
package main

import (
	"context"
	"log"

	"github.com/podhmo/ctxlog"
	"github.com/podhmo/ctxlog/zapctxlog"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	l := zapctxlog.MustNew()
	ctx, _ := ctxlog.WithLogger(context.Background(), l).With("x-id", 10)
	return f(ctx)
}

func f(ctx context.Context) error {
	ctx, log := ctxlog.Logger(ctx).With("y-id", 20)
	log.Info("start f")
	defer log.Info("end f")
	return g(ctx)
}

func g(ctx context.Context) error {
	ctx, log := ctxlog.Logger(ctx).With("y-id", 20)
	log.Info("start g")
	defer log.Info("end g")
	return nil
}
```

output

```console
$ go run _examples/00readme/main.go
{"level":"info","ts":1565388319.7551425,"caller":"00readme/main.go:29","msg":"start f","x-id":10,"y-id":20}
{"level":"info","ts":1565388319.7551734,"caller":"00readme/main.go:36","msg":"start g","x-id":10,"y-id":20,"y-id":20}
{"level":"info","ts":1565388319.7551854,"caller":"00readme/main.go:38","msg":"end g","x-id":10,"y-id":20,"y-id":20}
{"level":"info","ts":1565388319.755195,"caller":"00readme/main.go:31","msg":"end f","x-id":10,"y-id":20}
```

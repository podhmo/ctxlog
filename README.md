[![CircleCI](https://circleci.com/gh/podhmo/ctxlog.svg?style=svg)](https://circleci.com/gh/podhmo/ctxlog)

# ctxlog

contextual logging and structural logging

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
	ctx, _ := ctxlog.Set(context.Background(), l).With("x-id", 10)
	return f(ctx)
}

func f(ctx context.Context) error {
	ctx, log := ctxlog.Get(ctx).With("y-id", 20)
	log.Info("start f")
	defer log.Info("end f")
	return g(ctx)
}

func g(ctx context.Context) error {
	ctx, log := ctxlog.Get(ctx).With("y-id", 20)
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

### WithError()

```go
package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/podhmo/ctxlog"
	"github.com/podhmo/ctxlog/zapctxlog"
)


func main() {
	l := zapctxlog.MustNew()
	ctx, _ := ctxlog.Set(context.Background(), l).With("x-id", 10)
	if err := f(ctx); err != nil {
		ctxlog.Get(ctx).WithError(err).Warning("!!")
	}
}

func f(ctx context.Context) error {
	return errors.Wrap(fmt.Errorf("hmm"), "on f")
}
```

```
{
  "level":"warn",
  "ts":1565421996.0991826,
  "caller":"01readme/main.go:16",
  "msg":"!!",
  "x-id":10,
  "error":"on f: hmm",
  "errorVerbose":"hmm\non f\nmain.f\n\t$GOPATH/src/github.com/podhmo/ctxlog/zapctxlog/_examples/01readme/main.go:21\nmain.main\n\t$GOPATH/src/github.com/podhmo/ctxlog/zapctxlog/_examples/01readme/main.go:15\nruntime.main\n\t/usr/lib/go/src/runtime/proc.go:200\nruntime.goexit\n\t/usr/lib/go/src/runtime/asm_amd64.s:1337"
}
```

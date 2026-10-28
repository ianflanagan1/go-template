// Command app greets the name it's given. It's the composition root: the
// only place that builds adapters and passes them to the shell.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ianflanagan1/go-template/internal/adapters/clock"
	"github.com/ianflanagan1/go-template/internal/app"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "app:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 1 {
		return errors.New("usage: app <name>")
	}
	return app.Greet(clock.System{}, os.Stdout, args[0])
}

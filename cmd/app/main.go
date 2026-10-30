// Command app is the composition root: the only place that reads the
// configuration, builds adapters and passes them to the code that uses them.
package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err := run(ctx, os.Args[1:], os.Stdout, os.Stderr)
	stop()
	if err != nil {
		fmt.Fprintln(os.Stderr, "app:", err)
		os.Exit(1)
	}
}

// run reads the configuration from args, builds the adapters, passes them in
// and does the command's work, stopping when ctx is cancelled by an interrupt
// or SIGTERM. It writes output to stdout and logs to stderr. Name each
// parameter, (ctx context.Context, args []string, stdout, stderr io.Writer),
// as run starts to use it.
func run(_ context.Context, _ []string, _, _ io.Writer) error {
	return nil
}

// Command app is the composition root: the only place that builds adapters
// and passes them to the code that uses them.
package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "app:", err)
		os.Exit(1)
	}
}

// run builds the adapters, passes them in and does the command's work.
func run() error {
	return nil
}

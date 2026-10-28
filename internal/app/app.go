// Package app is the imperative shell: it gathers inputs through ports, calls
// the core and acts on what it returns.
package app

import (
	"fmt"
	"io"
	"time"

	"github.com/ianflanagan1/go-template/internal/core/greeting"
)

// Clock is the port for the current time.
type Clock interface {
	Now() time.Time
}

// Greet writes the greeting for name, for the time clock gives.
func Greet(clock Clock, out io.Writer, name string) error {
	if _, err := fmt.Fprintln(out, greeting.For(name, clock.Now().Hour())); err != nil {
		return fmt.Errorf("writing the greeting: %w", err)
	}
	return nil
}

// Package greeting decides how to greet someone. It's part of the functional
// core: pure functions over values, with no I/O or clock.
package greeting

import "fmt"

// For returns the greeting for name at the given hour of the day, 0 to 23.
func For(name string, hour int) string {
	switch {
	case hour < 12:
		return fmt.Sprintf("Good morning, %s.", name)
	case hour < 18:
		return fmt.Sprintf("Good afternoon, %s.", name)
	default:
		return fmt.Sprintf("Good evening, %s.", name)
	}
}

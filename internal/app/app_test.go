package app_test

import (
	"strings"
	"testing"
	"time"

	"github.com/ianflanagan1/go-template/internal/app"
)

type fixedClock struct{ now time.Time }

func (c fixedClock) Now() time.Time { return c.now }

func TestGreetWritesTheGreetingForTheClocksTime(t *testing.T) {
	var out strings.Builder
	evening := time.Date(2026, time.October, 27, 19, 0, 0, 0, time.UTC)
	if err := app.Greet(fixedClock{evening}, &out, "Ada"); err != nil {
		t.Fatal(err)
	}
	if got, want := out.String(), "Good evening, Ada.\n"; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

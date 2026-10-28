package greeting_test

import (
	"testing"

	"github.com/ianflanagan1/go-template/internal/core/greeting"
)

func TestTheGreetingFollowsThePartOfTheDay(t *testing.T) {
	for _, tc := range []struct {
		name string
		hour int
		want string
	}{
		{"just after midnight is morning", 0, "Good morning, Ada."},
		{"noon is afternoon", 12, "Good afternoon, Ada."},
		{"six in the evening is evening", 18, "Good evening, Ada."},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := greeting.For("Ada", tc.hour); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

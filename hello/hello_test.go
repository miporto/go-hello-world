package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("hello", func(t *testing.T) {
		got := Hello()
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

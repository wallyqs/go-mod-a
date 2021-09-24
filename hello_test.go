package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	expected := "Hello"
	if got != expected {
		t.Fatalf("Expected %v, got: %v", expected, got)
	}
}

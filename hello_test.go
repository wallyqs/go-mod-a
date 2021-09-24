package hello

import (
	"fmt"
	"testing"
	"github.com/wallyqs/go-mod-b"
)

func TestHello(t *testing.T) {
	got := Hello()
	expected := "Hello"
	if got != expected {
		t.Fatalf("Expected %v, got: %v", expected, got)
	}
}

func TestHelloWorld(t *testing.T) {
	got := fmt.Sprintf("%s %s", Hello(), world.World())
	expected := "Hello World"
	if got != expected {
		t.Fatalf("Expected %v, got: %v", expected, got)
	}
}

package hello

import (
	"fmt"
	"testing"
	"github.com/wallyqs/go-mod-a"
	"github.com/wallyqs/go-mod-b"
)

func TestHelloWorld(t *testing.T) {
	got := fmt.Sprintf("%s %s", hello.Hello(), world.World())
	expected := "Hello World"
	if got != expected {
		t.Fatalf("Expected %v, got: %v", expected, got)
	}
}

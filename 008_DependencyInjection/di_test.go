package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sam")

	got := buffer.String()
	want := "Hello, Sam"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

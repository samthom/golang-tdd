package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	times := 3
	repeated := Repeat("a", times)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("s", 2)
	fmt.Println(result)
	// Output: ss
}

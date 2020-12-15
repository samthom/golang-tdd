package arraynslices

import "testing"

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		assertCorrectMessage(t, got, want)
	})
}

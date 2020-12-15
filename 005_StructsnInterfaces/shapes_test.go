package structsninterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
func TestArea(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want float64, shape Shape) {
		t.Helper()
		if got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertCorrectMessage(t, got, want, shape)
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)

	})

	t.Run("Table driven test", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{12, 6}, 72.0},
			{Circle{10}, 314.1592653589793},
		}

		for _, tt := range areaTests {
			got := tt.shape.Area()
			assertCorrectMessage(t, got, tt.want, tt.shape)
		}
	})
}

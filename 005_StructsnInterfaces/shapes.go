package structsninterfaces

import "math"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Perimeter calculate perimeter of a shape using given values
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// Area calculate the area of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area calculate the area of the rectangle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

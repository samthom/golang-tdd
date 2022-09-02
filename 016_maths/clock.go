package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func SecondsInRadians(t time.Time) float64 {
	// return float64(t.Second()) * (math.Pi / 30)
    return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
    angle := SecondsInRadians(t)
    x := math.Sin(angle)
    y := math.Cos(angle)
    return Point{x,y}
}

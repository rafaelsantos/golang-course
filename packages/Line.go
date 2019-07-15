package main

import (
	"math"
)

//Uppercase: public (visible inside and outside package)
//Lowercase: private (visible only inside package)

//Point : represents a point
type Point struct {
	x float64
	y float64
}

//cathetus : private function
func cathetus(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return
}

//Distance : public function calculates the distance
func Distance(a, b Point) float64 {
	cx, cy := cathetus(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

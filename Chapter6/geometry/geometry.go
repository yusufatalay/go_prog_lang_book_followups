package geometry

import "math"

type Point struct{ X, Y float64 }

// traditional function
// calling this would be a function call "geometry.Distance"
func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// same thing, but as a method of the Point type
// p is this method's receiver
// calling this would be a method call  "Point.Distance"
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

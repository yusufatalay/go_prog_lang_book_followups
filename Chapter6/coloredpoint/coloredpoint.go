package main

import (
	"fmt"
	"image/color"
	"math"
)

type point struct{ X, Y float64 }

func (p point) distance(q point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}
func (p *point) scaleby(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type coloredpoint struct {
	point
	Color color.RGBA
}

func main() {
	var cp coloredpoint
	cp.X = 1
	fmt.Println(cp.point.X)
	cp.point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = coloredpoint{point{1, 1}, red}
	var q = coloredpoint{point{5, 4}, blue}
	fmt.Println(p.distance(q.point))
	p.scaleby(2)
	q.scaleby(2)
	fmt.Println(p.distance(q.point))
}

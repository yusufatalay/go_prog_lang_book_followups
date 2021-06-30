package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 300            // canvas size
	cells         = 100                 // number of cells
	xyrange       = 30.0                // axis ranges ( -xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (= 30degree)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin30 , cos30

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: white; stroke-width: 0.7' "+"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			// average z height for a cell
			avg_z := (az + bz + cz + dz) / 4

			if math.IsNaN(avg_z) {
				continue
			}
			color := "black"

			if avg_z > 0 {
				color = "red"
			} else {
				color = "blue"
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s;'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)

			//	fmt.Printf("az:%g  bz:%g  cz:%g  dz:%g   avg: %g\n\n", az, bz, cz, dz, avg_z)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx, sy, z float64) {
	// find point (x,y) at corner of cell (i,j)

	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute surfqce height z
	z := f(x, y)
	//fmt.Println(z)
	// project (x,y,z) isometrically ont o 2-d SVG canves (sx,sy)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return
}

func f(x, y float64) (ef float64) {

	r := math.Hypot(x, y) // distance from (0,0)
	ef = math.Sin(r) / r
	return
}

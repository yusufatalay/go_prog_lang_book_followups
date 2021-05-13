package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
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

	http.HandleFunc("/", handleRoot) // re-route the user to /design path and guide them to color and shape the mesh

	http.HandleFunc("/design", handleSVG) // design dir will draw the mesh accordinf to user's choice

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "to design the mesh go : /design?aboveQ=<color for z> 0.25>&aboveZ=<color>&belowQ=<color>&belowZ=<color>&meshHeight=<int>&meshWidth=<int>")
}

func handleSVG(w http.ResponseWriter, r *http.Request) {

	// handling the parameters like this order
	aboveQ := r.FormValue("aboveQ")
	aboveZ := r.FormValue("aboveZ")
	belowQ := r.FormValue("belowQ")
	belowZ := r.FormValue("belowZ")
	meshHeight := r.FormValue("meshHeight")
	meshWidth := r.FormValue("meshWidth")

	w.Header().Set("Content-Type", "image/svg+xml") // set the correct content type
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: white; stroke-width: 0.7' "+"width='%d' height='%d'>", meshWidth, meshHeight)

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

			if avg_z > 0.25 {
				color = aboveQ
			} else if avg_z > 0 {
				color = aboveZ
			} else if avg_z < -0.25 {
				color = belowQ
			} else if avg_z < 0 {
				color = belowZ
			} else {

				color = "black"
			}

			//			if avg_z > 0 {
			//				color = "red"
			//			} else {
			//				color = "blue"
			//			}

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s;'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)

		}
	}
	fmt.Fprintf(w, "</svg>")

}

func corner(i, j int) (float64, float64, float64) {
	// find point (x,y) at corner of cell (i,j)

	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute surfqce height z
	z := f(x, y)
	//fmt.Println(z)
	// project (x,y,z) isometrically ont o 2-d SVG canves (sx,sy)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {

	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

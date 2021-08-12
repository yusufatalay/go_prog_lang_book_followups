// Mandelbrot emits a PNG image of the Mandelbrot fractal
// implementing parallellism to accomplish Exercise 8.5
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var done = make(chan struct{})
	for i := 0; i < height; i++ {
		go func(i int) {
			for py := 0; py < height; py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// image point(px,py) represents complex value z.
					img.Set(px, py, mandelbrot(z))
				}
			}
			done <- struct{}{}
		}(i)
		png.Encode(os.Stdout, img) // NOT ignoring the errors
	}

	for i := 0; i < height; i++ {
		<-done
	}
}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const constrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - constrast*n}
		}
	}
	return color.Black
}

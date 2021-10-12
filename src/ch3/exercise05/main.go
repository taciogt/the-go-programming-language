// Mandebrot emits a PNG image of the Mandelbrot fractal
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		//width, height          = 1024, 1024
		width, height = 1024 * 10, 1024 * 10
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// IMage point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:

				g := 0xff * (float64(n) / float64(iterations))
				return color.RGBA{G: uint8(g), A: 0xff}
			default:
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{B: 255 - uint8(logScale*255), A: 255}
			}
		}
	}
	return color.Black
}

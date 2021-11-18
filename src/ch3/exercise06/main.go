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

		width, height = 1024 * 5, 1024 * 5
		epsX          = (xmax - xmin) / width
		epsY          = (ymax - ymin) / height
	)
	offX := []float64{-epsX, epsX}
	offY := []float64{-epsY, epsY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			//z := complex(x, y)

			// Supersampling
			subPixels := make([]color.Color, 0)
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					z := complex(x+offX[i], y+offY[j])
					subPixels = append(subPixels, mandelbrot(z))
				}
			}

			img.Set(px, py, avg(subPixels))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func avg(subPixels []color.Color) color.Color {
	total := uint32(len(subPixels))
	var r, g, b, a uint32
	for _, pixel := range subPixels {
		r_, g_, b_, a_ := pixel.RGBA()
		r += r_
		g += g_
		b += b_
		a += a_
	}
	return color.RGBA{
		R: uint8(r / total),
		G: uint8(g / total),
		B: uint8(b / total),
		A: uint8(a / total),
	}
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

// Surface computes an SVG rendering of a 3-D surface function
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30 degrees)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30dg), cos(30dg)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			color := getColor(i, j)
			if !math.IsNaN(ax + ay + bx + by + cx + cy + dx + dy) {
				fmt.Printf("<polygon fill='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n", color, ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x - y*cos30*xyscale)
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func getColor(i, j int) (color string) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	sin := math.Sin(math.Hypot(x, y))
	cs := fmt.Sprintf("%02x", int(math.Abs(sin*0xff)))
	if z > 0 {
		color = fmt.Sprintf("#0000%s", cs)
	} else {
		color = fmt.Sprintf("#%s0000", cs)
	}
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

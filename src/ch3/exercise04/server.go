// handler echoes the http request

// server up: go run . server.go &
// test: go run src/ch1/fetch/server.go http://localhost:8000/hello
// test: go run fetch/server.go http://localhost:8000/hello\?query\=123
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

var count int

func main() {
	http.HandleFunc("/", handler)       // each request call handler
	http.HandleFunc("/svg", svgHandler) // each request call handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q", k, v)
	}
}

func svgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	_ = r.ParseForm()
	wdt, err := strconv.Atoi(r.FormValue("width"))
	if err != nil {
		fmt.Printf("found error parsing width: %v", err)
		wdt = 600
	}
	hgt, err := strconv.Atoi(r.FormValue("height"))
	if err != nil {
		fmt.Printf("found error parsing height: %v", err)
		hgt = 320
	}
	setParams(wdt, hgt)

	writeSvg(w)
}

var (
	width, height = 600, 320                     // canvas size in pixels
	cells         = 100                          // number of grid cells
	xyrange       = 30.0                         // axis ranges (-xyrange..+xyrange)
	xyscale       = float64(width) / 2 / xyrange // pixels per x or y unit
	zscale        = float64(height) * 0.4        // pixels per z unit
	angle         = math.Pi / 6                  // angle of x, y axes (=30 degrees)
)

func setParams(w int, h int) {
	width, height = w, h                   // canvas size in pixels
	cells = 100                            // number of grid cells
	xyrange = 30.0                         // axis ranges (-xyrange..+xyrange)
	xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
	zscale = float64(height) * 0.4         // pixels per z unit
	angle = math.Pi / 6                    // angle of x, y axes (=30 degrees)

}

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30dg), cos(30dg)

func writeSvg(out io.Writer) {
	_, _ = fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
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
				_, _ = fmt.Fprintf(out, "<polygon fill='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n", color, ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	_, _ = fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z isometrically onto 2-D SVG canvas (sx, sy)
	sx := float64(width)/2 + (x - y*cos30*xyscale)
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func getColor(i, j int) (color string) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

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
